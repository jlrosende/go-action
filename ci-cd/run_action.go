package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := run_action(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func run_action(ctx context.Context) error {
	fmt.Println("Test binary installation")

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

	// get `node` image
	node := client.Container().From("node:16")

	// mount cloned repository into `node` image
	node = node.WithDirectory("/src", src).WithWorkdir("/src")

	// node = node.WithDirectory("/src", src).WithWorkdir("/src")
	node = node.WithEnvVariable("RUNNER_TEMP", "/temp/runner")
	node = node.WithEnvVariable("INPUT_SISU_VERSION", "0.0.1")

	node = node.WithExec([]string{"npm", "install"})

	node = node.WithExec([]string{"node", "setup-sisu.js"})

	node = node.WithExec([]string{"ls", "-la", "/temp/runner"})

	out, err := node.Stdout(ctx)
	if err != nil {
		return err
	}
	fmt.Println(out)
	for _, runner_os := range oses {
		for _, runner_arch := range arches {
			fmt.Println(runner_os, runner_arch)

		}
	}

	return nil
}
