package init

import (
	"embed"
	"errors"
	"os"
	"path"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	conf "github.com/jlrosende/go-action/config"
)

//go:embed templates/*
var tplDir embed.FS

var (
	iArgs   = InitArgs{}
	InitCmd = &cobra.Command{
		Use:   "init",
		Short: "",
		Long:  ``,
		Run:   initRepo,
	}
)

const (
	CLOUD_DIR     = "cloud"
	SISU_FILENAME = "sisu.yml"
	GITHUB_DIR    = ".github"
	WORKFLOWS_DIR = "workflows"
	ISSUES_DIR    = "ISSUE_TEMPLATE"
	TEMPLATES_DIR = "templates"
)

var (
	SISU_PATH      = path.Join(CLOUD_DIR, SISU_FILENAME)
	ISSUES_PATH    = path.Join(GITHUB_DIR, ISSUES_DIR)
	WORKFLOWS_PATH = path.Join(GITHUB_DIR, WORKFLOWS_DIR)
)

type InitArgs struct {
	Dir         string `json:"dir" yaml:"dir"`
	Overwrite   bool   `json:"overwrite" yaml:"overwrite"`
	Interactive bool   `json:"interactive" yaml:"interactive"`

	Runtime string `json:"runtime" yaml:"runtime"`
	Cloud   string `json:"cloud" yaml:"cloud"`
}

func init() {

	InitCmd.Flags().StringVarP(&iArgs.Dir, "directory", "d", ".", "Select the init directory (required)")
	InitCmd.Flags().BoolVar(&iArgs.Overwrite, "overwrite", false, "Select the init directory (required)")
	InitCmd.Flags().BoolVar(&iArgs.Interactive, "interactive", true, "Select the init directory (required)")

	InitCmd.Flags().StringVar(&iArgs.Runtime, "runtime", "", "Select the runtime")
	InitCmd.Flags().StringVar(&iArgs.Cloud, "cloud", "", "Select the cloud")

}

func initRepo(ccmd *cobra.Command, args []string) {

	_, err := os.Stat(CLOUD_DIR)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(CLOUD_DIR, 0755)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}

	_, err = os.Stat(SISU_PATH)
	if os.IsNotExist(err) {
		_, errFile := os.Create(SISU_PATH)
		if errFile != nil {
			log.Fatal(errFile)
		}
	}

	_, err = os.Stat(ISSUES_PATH)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(ISSUES_PATH, 0755)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}

	issue_files, err := tplDir.ReadDir(path.Join(TEMPLATES_DIR, "issues"))
	if err != nil {
		log.Errorf("can not read directory %v\n", err)
		return
	}

	for _, issue := range issue_files {
		src_issue_path := path.Join(TEMPLATES_DIR, "issues", issue.Name())

		fileContent, err := tplDir.ReadFile(src_issue_path)
		if err != nil {
			log.Errorf("can not read file %v\n", err)
			return
		}
		dest_issue_path := path.Join(ISSUES_PATH, issue.Name())
		err = os.WriteFile(dest_issue_path, fileContent, 0666)
		if err != nil {
			log.Errorf("can not write file %v\n", err)
			return
		}
		log.Debugf("Write file %s", dest_issue_path)
	}

	_, err = os.Stat(WORKFLOWS_PATH)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(WORKFLOWS_PATH, 0755)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}

	// var config conf.Config
	// tpls, err := template.New("templates").Delims("[[", "]]").ParseFS(templates, "templates/issues/*", "templates/workflows/*")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// log.Infof("%s", tpls.Execute(os.Stdout, []byte{}))

	// dirs, err := tplDir.ReadDir(TEMPLATES_DIR)
	// if err != nil {
	// 	log.Errorf("can not read directory %v\n", err)
	// 	return
	// }

	// for i, d := range dirs {
	// 	log.Infof("%d %s", i, d.Name())
	// }

	// var existConf bool = false

	// // Check if config already exist
	// log.Trace(viper.ConfigFileUsed())
	// if viper.ConfigFileUsed() != "" {
	// 	log.Trace("Config file already exist")
	// 	existConf = true
	// }

	// overwrite := false
	// if !iArgs.Overwrite && existConf {
	// 	log.Info("Sobreescribir")
	// 	overwrite = overwriteForm()
	// }

	// if iArgs.Interactive {
	// 	config = *form()
	// }

	// if iArgs.Overwrite || overwrite || !existConf {
	// 	saveConfig(config)
	// }

	// response, err := yaml.Marshal(iArgs)
	// if err != nil {
	// 	githubactions.Errorf("ERROR: %s", err.Error())
	// 	log.Fatal(err)
	// 	return
	// }
	// githubactions.SetOutput("args", string(response))

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

func defaultConf(cloud, runtime, region string) *conf.Config {
	return &conf.Config{
		Version: "0.0.0",
		Env: map[string][]conf.Function{
			"dev": {{
				Name:          "funtion-name-dev",
				Type:          "back",
				ResourceGroup: "resource-group-name",
				PackagePath:   "./code/dist/",
				Region:        region,
				Cloud:         cloud,
				Runtime:       runtime,
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
			"pre": {{
				Name:          "funtion-name-pre",
				Type:          "back",
				ResourceGroup: "resource-group-name",
				PackagePath:   "./code/dist/",
				Region:        region,
				Cloud:         cloud,
				Runtime:       runtime,
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
				Region:        region,
				Cloud:         cloud,
				Runtime:       runtime,
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
