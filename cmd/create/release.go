package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"r"},
	Short:   "",
	Long:    ``,
	RunE:    release,
}

func init() {

}

func release(ccmd *cobra.Command, args []string) error {
	fmt.Println("Create Relase")
	return nil
}
