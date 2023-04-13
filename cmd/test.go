package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Add git repository containing markdown content files",
		Long:  ``,
		Run:   test,
	}
)

func test(ccmd *cobra.Command, args []string) {
	fmt.Println("Deploy funtion")
}