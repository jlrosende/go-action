package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:     "snapshot",
	Aliases: []string{"s"},
	Short:   "",
	Long:    ``,
	RunE:    snapshot,
}

func init() {

}

func snapshot(ccmd *cobra.Command, args []string) error {
	fmt.Println("Create Snapshot")
	return nil
}
