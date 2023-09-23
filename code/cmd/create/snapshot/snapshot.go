package snapshot

import (
	"encoding/json"

	"regexp"

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

const VALID_TAG = "^[a-zA-Z0-9][a-zA-Z0-9_.-]{2,}$"

type CreateSnapshotArgs struct {
	Env          string `json:"env"`
	From         string `json:"from"`
	ForceRebuild bool   `json:"force-rebuild"`
	SnapshotName string `json:"snapshot-name"`
	Name         string `json:"name"`
	Cloud        string `json:"cloud"`
}

func init() {
	SnapshotCmd.Flags().StringVarP(&csArgs.Env, "environment", "e", "", "Select the environment (required)")
	SnapshotCmd.Flags().StringVarP(&csArgs.From, "from", "f", "", "Select the origin tag, branch or commit (required)")
	SnapshotCmd.Flags().StringVar(&csArgs.SnapshotName, "snapshot-name", "", "Tag name ^[a-zA-Z0-9][a-zA-Z0-9_.-]{2,}$ (required)")
	SnapshotCmd.Flags().StringVarP(&csArgs.Name, "name", "n", ".*", "Regex to select the matrix options from the list")
	SnapshotCmd.Flags().StringVarP(&csArgs.Cloud, "cloud", "c", ".*", "Regex to select which matrix options are select by cloud")

	SnapshotCmd.MarkFlagRequired("environment")
	SnapshotCmd.MarkFlagRequired("from")
	SnapshotCmd.MarkFlagRequired("snapshot-name")
	SnapshotCmd.MarkFlagsRequiredTogether("environment", "from", "snapshot-name")

	SnapshotCmd.Flags().BoolVar(&csArgs.ForceRebuild, "force-rebuild", false, "Force build and compilation of the artifact.")
}

func snapshot(cmd *cobra.Command, args []string) {

	r, err := regexp.Compile(VALID_TAG)
	if err != nil {
		githubactions.Errorf("ERROR: %s", err.Error())
		log.Fatal(err)
		return
	}

	if !r.MatchString(csArgs.SnapshotName) {
		githubactions.Errorf("Error Invalid Tag: %s ::Tag format: ^[a-zA-Z0-9][a-zA-Z0-9_.-]{2,}$", csArgs.SnapshotName)
		log.Fatalf("Error Invalid Tag: %s ::Tag format: ^[a-zA-Z0-9][a-zA-Z0-9_.-]{2,}$", csArgs.SnapshotName)
		return
	}

	response, err := json.Marshal(csArgs)
	if err != nil {
		githubactions.Errorf("ERROR: %s", err.Error())
		log.Fatal(err)
		return
	}
	githubactions.SetOutput("args", string(response))
}
