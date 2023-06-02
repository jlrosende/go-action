package create

import (
	"github.com/jlrosende/go-action/cmd/create/release"
	"github.com/jlrosende/go-action/cmd/create/snapshot"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "",
	Long:    ``,
}

func init() {
	CreateCmd.AddCommand(release.ReleaseCmd)
	CreateCmd.AddCommand(snapshot.SnapshotCmd)
}
