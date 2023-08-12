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

	// create empty directory to put build outputs
	outputs := client.Directory()

	// get `golang` image
	golang := client.Container().From("golang:1.20.6-bullseye")

	golang = golang.WithExec([]string{"apt", "update"})
	golang = golang.WithExec([]string{"apt", "install", "zip", "gzip", "tar", "-y"})

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")
	golang = golang.WithDirectory("/src/cmd/init/templates", templates)
	golang = golang.WithDirectory("/src/cmd/update/templates", templates)

	for _, goos := range oses {
		for _, goarch := range arches {
			// create a directory for each os and arch
			path := "build/"

			// set GOARCH and GOOS in the build environment
			build := golang.WithEnvVariable("GOOS", goos)
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
		}
	}
	// write build artifacts to host
	_, err = outputs.Export(ctx, ".")
	if err != nil {
		return err
	}

	return nil
}
