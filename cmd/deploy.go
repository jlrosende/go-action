package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"

	"github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	"github.com/jlrosende/go-action/types"
)

var (
	env       string
	from      string
	region    string
	name      string
	cloud     string
	deployCmd = &cobra.Command{
		Use:     "deploy",
		Aliases: []string{"d"},
		Short:   "",
		Long:    ``,
		RunE:    deploy,
	}
)

func init() {

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
	log.Tracef("Region: %s", region)
	log.Tracef("Name: %s", name)

	environs := config.Env

	if _, ok := environs[env]; !ok {
		env_values := reflect.ValueOf(environs).MapKeys()
		githubactions.Errorf("available environments are: %s", env_values)
		return fmt.Errorf("available environments are: %s", env_values)
	}
	// Filter by environment
	// options := viper.Get(fmt.Sprintf("environments.%s", env)).([]interface{})
	options := environs[env]
	log.Tracef("Number of functions in %s: %d:", env, len(options))

	// Filter by region
	filtered, err := filterByRegion(options, region)
	if err != nil {
		return err
	}

	// Filter by name
	filtered, err = filterByName(filtered, name)

	if err != nil {
		return err
	}

	log.Tracef("Number of functions to output: %d", len(filtered))

	response, err := json.Marshal(filtered)
	if err != nil {
		return err
	}

	githubactions.SetOutput("function_matrix", string(response))
	fmt.Println(string(response))

	return nil
}

func filterByRegion(options []types.Function, regex string) ([]types.Function, error) {

	filtered := []types.Function{}

	r, err := regexp.Compile(regex)
	if err != nil {
		return filtered, err
	}

	for i, e := range options {
		if r.MatchString(e.Region) {
			log.Tracef("[%d] %+v", i, e.Region)
			filtered = append(filtered, e)
		}
	}

	return filtered, nil
}

func filterByName(options []types.Function, regex string) ([]types.Function, error) {

	filtered := []types.Function{}

	r, err := regexp.Compile(regex)
	if err != nil {
		return filtered, err
	}

	for i, e := range options {
		if r.MatchString(e.Name) {
			log.Tracef("[%d] %+v", i, e.Name)
			filtered = append(filtered, e)
		}
	}

	return filtered, nil
}
