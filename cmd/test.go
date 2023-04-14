package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Add git repository containing markdown content files",
	Long:  ``,
	RunE:  test,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func test(ccmd *cobra.Command, args []string) error {
	fmt.Println("Deploy funtion")
	return nil
}
