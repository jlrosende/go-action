package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  ``,
	RunE:  initRepo,
}

func init() {

}

func initRepo(ccmd *cobra.Command, args []string) error {
	log.Info("Init")
	log.Info("Create sisu.yml")
	log.Info("Create actions")
	log.Info("-- Deploy AZF [0 - ]")
	log.Info("-- PR Int/Unit Test [1 - ]")
	log.Info("-- Create Release [2 - ]")
	log.Info("-- Create Snapshot [2 - ]")
	return nil
}
