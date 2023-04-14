package cmd

import (
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	releaseCmd = &cobra.Command{
		Use:   "release",
		Short: "Add git repository containing markdown content files",
		Long:  ``,
		RunE:  release,
	}
)

func init() {
	rootCmd.AddCommand(releaseCmd)
}

func release(cmd *cobra.Command, args []string) error {
	log.Info("RELEASE")
	return nil
}
