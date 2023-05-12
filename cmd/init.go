package cmd

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/sethvargo/go-githubactions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	conf "github.com/jlrosende/go-action/config"
)

var (
	iArgs   = InitArgs{}
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "",
		Long:  ``,
		Run:   initRepo,
		PreRun: func(cmd *cobra.Command, args []string) {

			if err := conf.LoadConfig(&config, cfgFile, iArgs.Dir); err != nil {
				return
			}
		},
	}
)

type InitArgs struct {
	Dir string `json:"dir" yaml:"dir"`
}

func init() {

	initCmd.Flags().StringVarP(&iArgs.Dir, "directory", "d", ".", "Select the init directory (required)")

}

func initRepo(ccmd *cobra.Command, args []string) {

	var existConf bool = false

	// Check if config already exist
	log.Trace(viper.ConfigFileUsed())
	if viper.ConfigFileUsed() != "" {
		log.Trace("Config file already exist")
		existConf = true
	}

	if existConf {
		log.Info("Sobreescribir")
	}

	// log.Info(dir)

	// log.Info("Init")
	// log.Info("Create sisu.yml")
	// log.Info("Create actions")
	// log.Info("-- Deploy AZF [0 - ]")
	// log.Info("-- PR Int/Unit Test [1 - ]")
	// log.Info("-- Create Release [2 - ]")
	// log.Info("-- Create Snapshot [2 - ]")

	response, err := yaml.Marshal(iArgs)
	if err != nil {
		githubactions.Errorf("ERROR: %s", err.Error())
		log.Fatal(err)
		return
	}
	githubactions.SetOutput("args", string(response))
}

func selectRuntime() string {
	items := []string{"java11", "java17", "node16", "node18"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select the function runtime",
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return ""
	}

	log.Tracef("You choose %s\n", result)
	return result
}

func selectCloud() string {
	items := []string{"aws", "azure", "gcp"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select the cloud",
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return ""
	}

	log.Tracef("You choose %s\n", result)
	return result
}

func selectRegion() string {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Username must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Region",
		Validate: validate,
		Default:  "eu-west-1",
	}

	result, err := prompt.Run()

	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return ""
	}

	log.Tracef("you select region %s\n", result)
	return result

}

func functionForm() conf.Function {
	return conf.Function{}
}

func environmentForm() map[string][]conf.Function {
	return map[string][]conf.Function{}
}
