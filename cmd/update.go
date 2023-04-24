package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  ``,
	RunE:  update,
}

func init() {

}

func update(ccmd *cobra.Command, args []string) error {
	fmt.Println("Update")
	return nil
}
