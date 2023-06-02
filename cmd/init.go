package cmd

import (
	"embed"
	"errors"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/sethvargo/go-githubactions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	conf "github.com/jlrosende/go-action/config"
)

//go:embed templates/*
var tplDir embed.FS

var (
	iArgs   = InitArgs{}
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "",
		Long:  ``,
		Run:   initRepo,
	}
)

type InitArgs struct {
	Dir         string `json:"dir" yaml:"dir"`
	Overwrite   bool   `json:"overwrite" yaml:"overwrite"`
	Interactive bool   `json:"interactive" yaml:"interactive"`
}

func init() {

	initCmd.Flags().StringVarP(&iArgs.Dir, "directory", "d", ".", "Select the init directory (required)")
	initCmd.Flags().BoolVar(&iArgs.Overwrite, "overwrite", false, "Select the init directory (required)")
	initCmd.Flags().BoolVar(&iArgs.Interactive, "interactive", true, "Select the init directory (required)")

}

func initRepo(ccmd *cobra.Command, args []string) {
	var config conf.Config
	// tpls, err := template.New("templates").Delims("[[", "]]").ParseFS(templates, "templates/issues/*", "templates/workflows/*")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// log.Infof("%s", tpls.Execute(os.Stdout, []byte{}))

	dirs, err := tplDir.ReadDir("templates")
	if err != nil {
		log.Errorf("can not read directory %v\n", err)
		return
	}

	for i, d := range dirs {
		log.Infof("%d %s", i, d.Name())
	}

	var existConf bool = false

	// Check if config already exist
	log.Trace(viper.ConfigFileUsed())
	if viper.ConfigFileUsed() != "" {
		log.Trace("Config file already exist")
		existConf = true
	}

	overwrite := false
	if !iArgs.Overwrite && existConf {
		log.Info("Sobreescribir")
		overwrite = overwriteForm()
	}

	if iArgs.Interactive {
		config = *form()
	}

	if iArgs.Overwrite || overwrite || !existConf {
		saveConfig(config)
	}

	response, err := yaml.Marshal(iArgs)
	if err != nil {
		githubactions.Errorf("ERROR: %s", err.Error())
		log.Fatal(err)
		return
	}
	githubactions.SetOutput("args", string(response))

	// log.Info(dir)
	// log.Info("Init")
	// log.Info("Create sisu.yml")
	// log.Info("Create actions")
	// log.Info("-- Deploy AZF [0 - ]")
	// log.Info("-- PR Int/Unit Test [1 - ]")
	// log.Info("-- Create Release [2 - ]")
	// log.Info("-- Create Snapshot [2 - ]")
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

func overwriteForm() bool {
	prompt := promptui.Select{
		Label: "Overwrite?[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func functionForm() conf.Function {
	return conf.Function{}
}

func environmentForm() map[string][]conf.Function {
	return map[string][]conf.Function{}
}

func saveConfig(config conf.Config) {
	data, err := yaml.Marshal(config)
	log.Tracef("\n%s", string(data))
	if err != nil {
		log.Fatal(err)
		return
	}

	fp, err := os.OpenFile("sisu.yaml", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fp.Close()
	fp.Write(data)
}

func form() *conf.Config {
	// example config
	return defaultConf()
	// new environment

	// new function

}

func defaultConf() *conf.Config {
	return &conf.Config{
		Version: rootCmd.Version,
		Env: map[string][]conf.Function{
			"dev": {
				{
					Name:          "funtion-name-dev",
					Type:          "back",
					ResourceGroup: "resource-group-name",
					PackagePath:   "./code/dist/",
					Region:        "westeurope",
					Cloud:         "azure",
					Runtime:       "java11",
					Database: &conf.Database{
						ResourceGroup: "resource-group-name",
						Name:          "db-name",
						Type:          "postgresql",
					},
					Vault: &conf.Vault{
						ResourceGroup: "resource-group-name",
						Name:          "vault-name",
					},
				},
			},
			"pre": {{
				Name:          "funtion-name-pre",
				Type:          "back",
				ResourceGroup: "resource-group-name",
				PackagePath:   "./code/dist/",
				Region:        "westeurope",
				Cloud:         "azure",
				Runtime:       "java11",
				Database: &conf.Database{
					ResourceGroup: "resource-group-name",
					Name:          "db-name",
					Type:          "postgresql",
				},
				Vault: &conf.Vault{
					ResourceGroup: "resource-group-name",
					Name:          "vault-name",
				},
			}},
			"pro": {{
				Name:          "funtion-name-pro",
				Type:          "back",
				ResourceGroup: "resource-group-name",
				PackagePath:   "./code/dist/",
				Region:        "westeurope",
				Cloud:         "azure",
				Runtime:       "java11",
				Database: &conf.Database{
					ResourceGroup: "resource-group-name",
					Name:          "db-name",
					Type:          "postgresql",
				},
				Vault: &conf.Vault{
					ResourceGroup: "resource-group-name",
					Name:          "vault-name",
				},
			}},
		},
	}
}
