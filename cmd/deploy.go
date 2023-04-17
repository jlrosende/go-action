package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"

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
	// environs := viper.GetStringMap("environments")

	environs := config.Env

	log.Tracef("%+v", environs)

	if _, ok := environs[env]; !ok {
		env_values := reflect.ValueOf(environs).MapKeys()
		githubactions.Errorf("available environments are: %s", env_values)
		return fmt.Errorf("available environments are: %s", env_values)
	}
	// Filter by environment
	// options := viper.Get(fmt.Sprintf("environments.%s", env)).([]interface{})
	options := environs[env]
	log.Tracef("Number of functions in %s: %d:", env, len(options))

	/*
		TODO Filter by region
		// filtered, err := filterBy(options, region)
		// if err != nil {
		// 	return err
		// }
	*/

	/*
		TODO Filter by name
		// filtered, err = filterBy(filtered, name)

		// if err != nil {
		// 	return err
		// }
	*/

	log.Tracef("Number of functions to output: %d", len(options))

	response, err := json.MarshalIndent(options, "", "  ")
	if err != nil {
		return err
	}
	log.Infof("function_matrix output: \n%s", string(response))

	response, err = json.Marshal(options)
	if err != nil {
		return err
	}

	githubactions.SetOutput("function_matrix", string(response))

	return nil
}

/*
func filterBy(options []types.Function, regex string) ([]types.Function, error) {

	// filtered := []types.Function{}

	// r, err := regexp.Compile(region)
	// if err != nil {
	// 	return filtered, err
	// }

	// for i, e := range options {
	// 	if r.MatchString(e.(map[string]interface{})["region"].(string)) {
	// 		log.Tracef("[%d] %+v", i, e.(map[string]interface{})["region"])
	// 		filtered = append(filtered, e)
	// 	}
	// }

	return nil, errors.New("unimplemented method")
}
*/
