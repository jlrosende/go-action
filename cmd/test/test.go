package test

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Add git repository containing markdown content files",
	Long:  ``,
	RunE:  test,
}

func init() {

}

func test(ccmd *cobra.Command, args []string) error {
	fmt.Println("Test")
	return nil
}
