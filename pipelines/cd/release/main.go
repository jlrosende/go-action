package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
)

var tag string

func init() {
	flag.StringVar(&tag, "tag", "", "Set tag of the curret build")
	flag.Parse()
}

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// define build matrix
	oses := []string{"linux", "darwin", "windows"}
	arches := []string{"amd64", "arm64"}

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory(".")

	template_path, err := filepath.Abs("./templates")
	if err != nil {
		return err
	}
	fmt.Println(template_path)

	templates := client.Host().Directory(template_path)

	goBuildCache := client.Host().Directory(os.Getenv("GO_CACHE"))
	goModCache := client.Host().Directory(os.Getenv("GO_MODCACHE"))

	fmt.Println(os.Getenv("GO_CACHE"))
	fmt.Println(os.Getenv("GO_MODCACHE"))
	// return errors.New("Set new caches")

	// create empty directory to put build outputs
	outputs := client.Directory()

	// get `golang` image
	golang := client.Container().From("golang:1.21-bookworm")

	golang = golang.WithDirectory("/root/.cache/go-build", goBuildCache)
	golang = golang.WithDirectory("/go/pkg/mod", goModCache)

	golang = golang.WithExec([]string{"apt", "update"})
	golang = golang.WithExec([]string{"apt", "install", "zip", "gzip", "tar", "-y"})

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")
	golang = golang.WithDirectory("/src/cmd/init/templates", templates)
	golang = golang.WithDirectory("/src/cmd/update/templates", templates)

	golang = golang.WithExec([]string{"go", "mod", "download"})

	var build *dagger.Container
	for _, goos := range oses {
		for _, goarch := range arches {
			// create a directory for each os and arch
			path := "build/"

			// set GOARCH and GOOS in the build environment
			build = golang.WithEnvVariable("GOOS", goos)
			build = build.WithEnvVariable("GOARCH", goarch)
			build = build.WithEnvVariable("CGO_ENABLED", "0")

			var fileName string
			if tag == "" {
				fileName = fmt.Sprintf("sisu-%s-%s", goos, goarch)
			} else {
				fileName = fmt.Sprintf("sisu-%s-%s-%s", tag, goos, goarch)
			}
			filePath := filepath.Join(path, fileName)

			// build application
			build = build.WithExec([]string{"go", "build", "-o", filePath, fmt.Sprintf("-ldflags=-X 'github.com/jlrosende/go-action/config.Version=%s'", tag)})

			if goos == "windows" {
				build = build.WithExec([]string{"zip", "-r", "-j", fmt.Sprintf("%s.zip", filePath), filePath})
			} else {
				build = build.WithExec([]string{"tar", "-C", "./build", "-czvf", fmt.Sprintf("%s.tar.gz", filePath), fileName})
			}

			build = build.WithExec([]string{"rm", filePath})

			// get reference to build output directory in container
			outputs = outputs.WithDirectory(path, build.Directory(path))
			goBuildCache = goBuildCache.WithDirectory(os.Getenv("GO_CACHE"), build.Directory("/root/.cache/go-build"))
			goModCache = goModCache.WithDirectory(os.Getenv("GO_MODCACHE"), build.Directory("/go/pkg/mod"))
		}
	}
	// write build artifacts to host
	_, err = outputs.Export(ctx, ".")
	if err != nil {
		return err
	}

	_, err = goBuildCache.Export(ctx, os.Getenv("GO_CACHE"))
	if err != nil {
		return err
	}
	_, err = goModCache.Export(ctx, os.Getenv("GO_MODCACHE"))
	if err != nil {
		return err
	}

	return nil
}
