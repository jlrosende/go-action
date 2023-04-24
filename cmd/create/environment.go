package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"env", "e"},
	Short:   "",
	Long:    ``,
	RunE:    environment,
}

func init() {

}

func environment(ccmd *cobra.Command, args []string) error {
	fmt.Println("Create Environment")
	return nil
}
