package snapshot

import (
	"encoding/json"

	"github.com/sethvargo/go-githubactions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	csArgs      = CreateSnapshotArgs{}
	SnapshotCmd = &cobra.Command{
		Use:     "snapshot",
		Aliases: []string{"s"},
		Short:   "",
		Long:    ``,
		Run:     snapshot,
	}
)

type CreateSnapshotArgs struct {
	Env          string `json:"env"`
	From         string `json:"from"`
	To           string `json:"to"`
	ForceRebuild bool   `json:"force-rebuild"`
}

func init() {
	SnapshotCmd.Flags().StringVarP(&csArgs.Env, "environment", "e", "", "Select the environment (required)")
	SnapshotCmd.Flags().StringVarP(&csArgs.From, "from", "f", "", "Select the origin tag, branch or commit (required)")
	SnapshotCmd.Flags().StringVarP(&csArgs.To, "to", "t", "", "Select the destination branch (required)")
	SnapshotCmd.MarkFlagRequired("environment")
	SnapshotCmd.MarkFlagRequired("from")
	SnapshotCmd.MarkFlagsRequiredTogether("environment", "from", "to")

	SnapshotCmd.Flags().BoolVar(&csArgs.ForceRebuild, "force-rebuild", false, "Force build and compilation of the artifact.")
}

func snapshot(cmd *cobra.Command, args []string) {
	response, err := json.Marshal(csArgs)
	if err != nil {
		githubactions.Errorf("ERROR: %s", err.Error())
		log.Fatal(err)
		return
	}
	githubactions.SetOutput("args", string(response))
}
