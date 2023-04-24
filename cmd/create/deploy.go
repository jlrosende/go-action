package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:     "deployment",
	Aliases: []string{"deploy", "d"},
	Short:   "",
	Long:    ``,
	RunE:    deploy,
}

func init() {

}

func deploy(ccmd *cobra.Command, args []string) error {
	fmt.Println("Create Deploy")
	return nil
}
