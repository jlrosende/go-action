package create

import (
	"context"
	"encoding/json"

	"github.com/google/go-github/github"
	"github.com/sethvargo/go-githubactions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	crArgs     = CreateReleaseArgs{}
	token      string
	releaseCmd = &cobra.Command{
		Use:     "release",
		Aliases: []string{"r"},
		Short:   "",
		Long:    ``,
		Run:     release,
	}
)

type CreateReleaseArgs struct {
	Env       string `json:"env"`
	From      string `json:"from"`
	Version   string `json:"version,omitempty"`
	Increment string `json:"increment,omitempty"`
}

func init() {
	releaseCmd.Flags().StringVarP(&crArgs.Env, "environment", "e", "", "Select the environment to be matrixed (required)")
	releaseCmd.Flags().StringVarP(&crArgs.From, "from", "f", "", "Select the tag, branch or commit to be matrixed (required)")
	releaseCmd.MarkFlagRequired("environment")
	releaseCmd.MarkFlagRequired("from")
	releaseCmd.MarkFlagsRequiredTogether("environment", "from")

	releaseCmd.Flags().StringVarP(&crArgs.Version, "version", "v", "", "TODO")
	releaseCmd.Flags().StringVarP(&crArgs.Increment, "increment", "i", "patch", "TODO")

	releaseCmd.MarkFlagsMutuallyExclusive("version", "increment")

	token = githubactions.GetInput("token")

	if token == "" {
		githubactions.Fatalf("missing 'token'")
	}
}

func release(ccmd *cobra.Command, args []string) {

	client := github.NewClient(nil)

	ghctx, err := githubactions.Context()
	if err != nil {
		log.Error(err)
		return
	}

	repo, _, err := client.Repositories.Get(context.Background(), ghctx.RepositoryOwner, ghctx.Repository)
	if err != nil {
		log.Error(err)
		return
	}

	log.Trace(repo)

	response, err := json.Marshal(crArgs)
	if err != nil {
		log.Error(err)
		return
	}
	githubactions.SetOutput("args", string(response))

}
