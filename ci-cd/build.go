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

	// create empty directory to put build outputs
	outputs := client.Directory()

	// get `golang` image
	golang := client.Container().From("golang:latest")

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

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
				fileName = fmt.Sprintf("go-action-%s-%s", goos, goarch)
			} else {
				fileName = fmt.Sprintf("go-action-%s-%s-%s", tag, goos, goarch)	
			}
			filePath := filepath.Join(path, fileName)

			// build application
			build = build.WithExec([]string{"go", "build", "-o", filePath})

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
