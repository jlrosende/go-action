package init

import (
	"bytes"
	"embed"
	"errors"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/sethvargo/go-githubactions"
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
	Runtime   string `json:"runtime" yaml:"runtime"`
	Cloud     string `json:"cloud" yaml:"cloud"`
	Region    string `json:"region" yaml:"region"`
	Overwrite bool   `json:"overwrite" yaml:"overwrite"`
}

func init() {

	InitCmd.Flags().StringVar(&iArgs.Runtime, "runtime", "", "Select the runtime")
	InitCmd.Flags().StringVar(&iArgs.Cloud, "cloud", "", "Select the cloud")
	InitCmd.Flags().StringVar(&iArgs.Region, "region", "", "Select the region")
	InitCmd.Flags().BoolVar(&iArgs.Overwrite, "overwrite", false, "Overwrite existing files")

}

func initRepo(cmd *cobra.Command, args []string) {

	_, err := os.Stat(ISSUES_PATH)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(ISSUES_PATH, 0755)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}

	err = DirectoryFromTemplate(ISSUES_PATH, path.Join(TEMPLATES_DIR, "issues"))
	if err != nil {
		log.Errorf("ISSUES: %v\n", err)
		return
	}

	if iArgs.Runtime == "" {
		iArgs.Runtime = selectRuntime()
	}

	var lang string = ""
	switch strings.ToLower(iArgs.Runtime) {
	case "java11", "java17":
		lang = "java"
	case "node16", "node18":
		lang = "node"
	case "go119", "go120":
		lang = "go"
	default:
		lang = "java"
	}

	err = DirectoryFromTemplate(WORKFLOWS_PATH, path.Join(TEMPLATES_DIR, "workflows", lang))
	if err != nil {
		log.Errorf("WORKFLOWS: %v\n", err)
		return
	}

	if iArgs.Cloud == "" {
		iArgs.Cloud = selectCloud()
	}

	if iArgs.Region == "" {
		iArgs.Region = selectRegion()
	}

	config := defaultConf(iArgs.Cloud, iArgs.Runtime, iArgs.Region)

	_, err = os.Stat(CLOUD_DIR)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(CLOUD_DIR, 0755)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}

	_, err = os.Stat(SISU_PATH)

	if !os.IsNotExist(err) && !iArgs.Overwrite {
		iArgs.Overwrite = askOverwrite()
	}

	if iArgs.Overwrite {
		err := saveConfig(SISU_PATH, config)
		if err != nil {
			log.Errorf("SISU CONF: %v\n", err)
			return
		}
	}

	response, err := yaml.Marshal(iArgs)
	if err != nil {
		githubactions.Errorf("ERROR: %s", err.Error())
		log.Errorf("GITHUB_OUTPUT: %v\n", err)
		return
	}
	githubactions.SetOutput("args", string(response))
}

func askOverwrite() bool {
	prompt := promptui.Prompt{
		Label:     "Overwrite files",
		IsConfirm: true,
		Default:   "y",
	}

	result, err := prompt.Run()

	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return false
	}

	log.Tracef("You choose %s\n", result)
	return true
}

func selectRuntime() string {
	items := []string{
		"java11",
		"java17",
		"node16",
		"node18",
		"g119",
		"go120",
	}
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
		if err != nil {
			log.Panic(err)
		}

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
			return errors.New("username must have more than 3 characters")
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

func saveConfig(filePath string, config *conf.Config) error {
	var b bytes.Buffer
	ymlEncoder := yaml.NewEncoder(&b)
	ymlEncoder.SetIndent(2)
	ymlEncoder.Encode(config)

	fp, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	fp.Write(b.Bytes())
	return nil
}

func DirectoryFromTemplate(directory, template_dir string) error {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}

	tmpl_files, err := tplDir.ReadDir(template_dir)
	if err != nil {
		return err
	}

	for _, file := range tmpl_files {
		src_tmpl_path := path.Join(template_dir, file.Name())

		tmpl, err := template.New(file.Name()).Delims("<<", ">>").ParseFS(tplDir, src_tmpl_path)
		if err != nil {
			return err
		}

		dest_tmpl_path := path.Join(directory, file.Name())

		fp, err := os.Create(dest_tmpl_path)
		if err != nil {
			return err
		}

		err = tmpl.Execute(fp, struct {
			Version string
		}{
			Version: conf.Version,
		})
		if err != nil {
			return err
		}

		// fileContent, err := tplDir.ReadFile(src_tmpl_path)
		// if err != nil {
		// 	return err
		// }

		// dest_tmpl_path := path.Join(directory, file.Name())
		// err = os.WriteFile(dest_tmpl_path, fileContent, 0666)
		// if err != nil {
		// 	return err
		// }

		log.Debugf("Write file %s", dest_tmpl_path)
	}
	return nil
}

func defaultConf(cloud, runtime, region string) *conf.Config {
	return &conf.Config{
		Version: conf.Version,
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
				Swap: conf.Swap{
					Mode:        "slot",
					FrontDoor:   nil,
					AppInsights: nil,
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
				Swap: conf.Swap{
					Mode:        "slot",
					FrontDoor:   nil,
					AppInsights: nil,
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
				Swap: conf.Swap{
					Mode:        "slot",
					FrontDoor:   nil,
					AppInsights: nil,
				},
			}},
		},
	}
}
