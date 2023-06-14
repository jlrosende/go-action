package matrix

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"

	"github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	conf "github.com/jlrosende/go-action/config"
)

var (
	cfgFile   string
	config    *conf.Config
	mArgs     = MatrixArgs{}
	MatrixCmd = &cobra.Command{
		Use:     "matrix",
		Aliases: []string{"deploy", "d", "m"},
		Short:   "Create a matrix of functions",
		Long:    `Create a matrix of functions filtered with the command args.`,
		Run:     matrix,
	}
)

type MatrixArgs struct {
	Env    string `json:"env"`
	From   string `json:"from"`
	Region string `json:"region,omitempty"`
	Name   string `json:"name,omitempty"`
	Cloud  string `json:"cloud,omitempty"`
	RunBD bool   `json:"run-db"`
}

func init() {

	MatrixCmd.Flags().StringVarP(&mArgs.Env, "environment", "e", "", "Select the environment to be matrixed (required)")
	MatrixCmd.Flags().StringVarP(&mArgs.From, "from", "f", "", "Select the tag, branch or commit to be matrixed (required)")
	MatrixCmd.MarkFlagRequired("environment")
	MatrixCmd.MarkFlagRequired("from")
	MatrixCmd.MarkFlagsRequiredTogether("environment", "from")

	// Filter by name default * to select all options
	MatrixCmd.Flags().StringVarP(&mArgs.Name, "name", "n", ".*", "Regex to select the matrix options from the list")

	// Filter by region
	MatrixCmd.Flags().StringVarP(&mArgs.Region, "region", "r", ".*", "Regex to select which matrix options are select by region")

	// Filter by cloud
	MatrixCmd.Flags().StringVarP(&mArgs.Cloud, "cloud", "c", ".*", "Regex to select which matrix options are select by cloud")

	MatrixCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is sisu.{yml,yaml})")
	
	MatrixCmd.Flags().BoolVar(&mArgs.RunBD, "run-db", true, "Run the action to deploy database configuration")

}

func matrix(cmd *cobra.Command, args []string) {

	log.Trace("cfgFile", cfgFile)
	config, err := conf.LoadConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Tracef("Environment: %+v\n", mArgs.Env)
	log.Tracef("From: %+v\n", mArgs.From)
	log.Tracef("Region: %s", mArgs.Region)
	log.Tracef("Name: %s", mArgs.Name)

	environs := config.Env

	if _, ok := environs[mArgs.Env]; !ok {
		env_values := reflect.ValueOf(environs).MapKeys()
		githubactions.Errorf("available environments are: %s", env_values)
		log.Errorf("available environments are: %s", env_values)
		return
	}
	// Filter by environment
	// options := viper.Get(fmt.Sprintf("environments.%s", env)).([]interface{})
	options := environs[mArgs.Env]
	log.Tracef("Number of functions in %s: %d:", mArgs.Env, len(options))

	// Filter by region
	filtered, err := filterByRegion(options, mArgs.Region)
	if err != nil {
		log.Error(err)
		return
	}

	// Filter by name
	filtered, err = filterByName(filtered, mArgs.Name)

	if err != nil {
		log.Error(err)
		return
	}

	log.Tracef("Number of functions to output: %d", len(filtered))

	response, err := json.Marshal(filtered)
	if err != nil {
		log.Error(err)
		return
	}

	githubactions.SetOutput("function_matrix", string(response))
	fmt.Println(string(response))

	response, err = json.Marshal(mArgs)
	if err != nil {
		log.Error(err)
		return
	}
	githubactions.SetOutput("args", string(response))

}

func filterByRegion(options []conf.Function, regex string) ([]conf.Function, error) {

	filtered := []conf.Function{}

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

func filterByName(options []conf.Function, regex string) ([]conf.Function, error) {

	filtered := []conf.Function{}

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
