package create

import (
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "",
	Long:    ``,
}

func init() {
	CreateCmd.AddCommand(releaseCmd)
	CreateCmd.AddCommand(snapshotCmd)
}
