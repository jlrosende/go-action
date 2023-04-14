package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"

	"github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var (
	env       string
	from      string
	region    string
	name      string
	cloud     string
	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Add git repository containing markdown content files",
		Long:  ``,
		RunE:  deploy,
	}
)

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&env, "environment", "e", "", "Select the environment to be deployed (required)")
	deployCmd.Flags().StringVarP(&from, "from", "f", "", "Select the tag, branch or commit to be deployed (required)")
	deployCmd.MarkFlagRequired("environment")
	deployCmd.MarkFlagRequired("from")

	deployCmd.MarkFlagsRequiredTogether("environment", "from")
	// Filter by name default * to select all options
	deployCmd.Flags().StringVarP(&name, "name", "n", ".*", "Regex to select the deployment options from the list")

	// Filter by region
	deployCmd.Flags().StringVarP(&region, "region", "r", ".*", "Regex to select which deployment options are select by region")
	// Filter by cloud

	deployCmd.Flags().StringVarP(&cloud, "cloud", "c", ".*", "Regex to select which deployment options are select by cloud")
}

func deploy(cmd *cobra.Command, args []string) error {
	log.Tracef("Environment: %+v\n", env)
	log.Tracef("From: %+v\n", from)

	// Validate environment
	environs := viper.GetStringMap("environments")

	log.Tracef("%+v", environs)

	if _, ok := environs[env]; !ok {
		env_values := reflect.ValueOf(environs).MapKeys()
		githubactions.Errorf("available environments are: %s", env_values)
		return fmt.Errorf("available environments are: %s", env_values)
	}
	// Filter by environment
	options := viper.Get(fmt.Sprintf("environments.%s", env)).([]interface{})
	log.Tracef("N_Options: %+v", len(options))

	// Filter by region
	filtered, err := filterBy(options, region)
	if err != nil {
		return err
	}

	// Filter by name
	filtered, err = filterBy(filtered, name)

	if err != nil {
		return err
	}

	log.Debug(len(filtered))

	response, err := json.Marshal(filtered)
	if err != nil {
		return err
	}

	log.Debug(string(response))

	githubactions.SetOutput("hola", string(response))

	return nil
}

func filterBy(options []interface{}, regex string) ([]interface{}, error) {

	filtered := []interface{}{}

	r, err := regexp.Compile(region)
	if err != nil {
		return filtered, err
	}

	for i, e := range options {
		if r.MatchString(e.(map[string]interface{})["region"].(string)) {
			log.Tracef("[%d] %+v", i, e.(map[string]interface{})["region"])
			filtered = append(filtered, e)
		}
	}

	return filtered, nil
}
