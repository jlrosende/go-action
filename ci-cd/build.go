package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"context"
	"flag"
	"fmt"
	"io"
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
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory(".")

	// create empty directory to put build outputs
	outputs := client.Directory()

	// get `golang` image
	golang := client.Container().From("golang:1.20.6-bullseye")

	golang = golang.WithExec([]string{"apt", "update"})
	golang = golang.WithExec([]string{"apt", "install", "zip", "gzip", "tar", "-y"})

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

func compress(file, osType string) error {
	if osType == "windows" {
		archive, err := os.Create(fmt.Sprintf("%s.zip", file))
		if err != nil {
			return err
		}
		defer archive.Close()

		//Create a new zip writer
		zipWriter := zip.NewWriter(archive)
		fmt.Println("opening first file")
		//Add files to the zip archive
		f1, err := os.Open("file")
		if err != nil {
			return err
		}
		defer f1.Close()

		fmt.Println("adding file to archive..")
		w1, err := zipWriter.Create("file")
		if err != nil {
			return err
		}
		if _, err := io.Copy(w1, f1); err != nil {
			return err
		}
		fmt.Println("closing archive")
		defer zipWriter.Close()
		return nil
	} else {
		archive, err := os.Create(fmt.Sprintf("%s.tar.gz", file))
		if err != nil {
			return err
		}
		defer archive.Close()

		gw := gzip.NewWriter(archive)
		defer gw.Close()
		tarWriter := tar.NewWriter(gw)
		defer tarWriter.Close()

		// Open the file which will be written into the archive
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		// Get FileInfo about our file providing file size, mode, etc.
		info, err := f.Stat()
		if err != nil {
			return err
		}

		// Create a tar Header from the FileInfo data
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// Use full path as name (FileInfoHeader only takes the basename)
		// If we don't do this the directory strucuture would
		// not be preserved
		// https://golang.org/src/archive/tar/common.go?#L626
		header.Name = file

		// Write file header to the tar archive
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}

		// Copy file content to tar archive
		_, err = io.Copy(tarWriter, f)
		if err != nil {
			return err
		}

		return nil
	}
}
