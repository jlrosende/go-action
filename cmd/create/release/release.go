package release

import (
	"encoding/json"

	"github.com/jlrosende/go-action/utils"
	"github.com/sethvargo/go-githubactions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	crArgs     = CreateReleaseArgs{}
	token      string
	ReleaseCmd = &cobra.Command{
		Use:     "release",
		Aliases: []string{"r"},
		Short:   "",
		Long:    ``,
		Run:     release,
	}
)

type CreateReleaseArgs struct {
	Env          string `json:"env"`
	From         string `json:"from"`
	Version      string `json:"version"`
	Increment    string `json:"increment"`
	ForceRebuild bool   `json:"force-rebuild"`
}

func init() {

	ReleaseCmd.Flags().StringVarP(&crArgs.Env, "environment", "e", "", "Select the environment (required)")
	ReleaseCmd.Flags().StringVarP(&crArgs.From, "from", "f", "", "Select the tag, branch or commit (required)")
	ReleaseCmd.MarkFlagRequired("environment")
	ReleaseCmd.MarkFlagRequired("from")
	ReleaseCmd.MarkFlagsRequiredTogether("environment", "from")

	ReleaseCmd.Flags().StringVarP(&crArgs.Version, "version", "v", "", "Manualy set the version, disable autoincrement")
	ReleaseCmd.Flags().StringVarP(&crArgs.Increment, "increment", "i", "patch", "Select which part of version is autoincremented (patch, minor, major) (default: patch)")
	ReleaseCmd.MarkFlagsMutuallyExclusive("version", "increment")

	ReleaseCmd.Flags().BoolVar(&crArgs.ForceRebuild, "force-rebuild", false, "Force build and compilation of the artifact.")

}

func release(cmd *cobra.Command, args []string) {

	// Validate args
	bumpOps := []string{"patch", "minor", "major"}
	if !utils.Contains(bumpOps, crArgs.Increment) {
		githubactions.Errorf("ERROR: --increment only allow the following arguments %s", bumpOps)
		log.Fatalf("ERROR: --increment only allow the following arguments %s", bumpOps)
		return
	}

	response, err := json.Marshal(crArgs)
	if err != nil {
		githubactions.Errorf("ERROR: %s", err.Error())
		log.Fatal(err)
		return
	}
	githubactions.SetOutput("args", string(response))

}
