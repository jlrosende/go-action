package update

import (
	"embed"
	"os"
	"path"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sethvargo/go-githubactions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	conf "github.com/jlrosende/go-action/config"
	"github.com/jlrosende/go-action/utils"
)

//go:embed templates/*
var tplDir embed.FS

var (
	cfgFile   string
	UpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "",
		Long:  ``,
		Run:   update,
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
	ISSUES_PATH    = path.Join(GITHUB_DIR, ISSUES_DIR)
	WORKFLOWS_PATH = path.Join(GITHUB_DIR, WORKFLOWS_DIR)
)

type Workflow struct {
	Env Env `yaml:"env"`
}

type Env struct {
	WORKFLOW_VERSION string `yaml:"WORKFLOW_VERSION"`
	UPDATE_WORKFLOW  bool   `yaml:"UPDATE_WORKFLOW"`
}

func init() {
	UpdateCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is sisu.{yml,yaml})")
}

func update(cmd *cobra.Command, args []string) {

	log.Trace("cfgFile", cfgFile)
	config, err := conf.LoadConfig(cfgFile)
	if err != nil {
		switch e := err.(type) {
		case validator.ValidationErrors:
			for _, errType := range e {
				githubactions.Errorf("ERROR: %s", errType)
				log.Errorf("ERROR: %s", errType)
			}
			return
		default:
			githubactions.Errorf("ERROR: %s", err)
			log.Error(err)
			return
		}
	}

	lang := config.GetLang()

	// Obtener los workflows del repo
	repo_workflows, err := os.ReadDir(WORKFLOWS_PATH)
	if err != nil {
		log.Fatal(err)
	}

	sisu_workflows := []string{}
	for _, archivo := range repo_workflows {

		if !strings.HasPrefix(archivo.Name(), "sisu-") {
			continue
		}

		// log.Debug(path.Join(WORKFLOWS_PATH, archivo.Name()))
		body, err := os.ReadFile(path.Join(WORKFLOWS_PATH, archivo.Name()))
		if err != nil {
			log.Error(err)
			return
		}
		var workflow Workflow
		err = yaml.Unmarshal(body, &workflow)
		if err != nil {
			log.Error(err)
			return
		}

		if !workflow.Env.UPDATE_WORKFLOW || workflow.Env.WORKFLOW_VERSION == conf.Version {
			sisu_workflows = append(sisu_workflows, archivo.Name())
			continue
		}
	}

	log.Debug(sisu_workflows)

	// Obtener los workflows del template
	tmpl_files, err := tplDir.ReadDir(path.Join(TEMPLATES_DIR, "workflows", lang))
	if err != nil {
		log.Fatal(err)
	}

	for _, archivo := range tmpl_files {
		log.Debug(archivo.Name())

		if utils.Contains(sisu_workflows, archivo.Name()) {
			continue
		}

		src_tmpl_path := path.Join(path.Join(TEMPLATES_DIR, "workflows", lang), archivo.Name())

		fileContent, err := tplDir.ReadFile(src_tmpl_path)
		if err != nil {
			log.Error(err)
			return
		}
		dest_tmpl_path := path.Join(WORKFLOWS_PATH, archivo.Name())
		err = os.WriteFile(dest_tmpl_path, fileContent, 0644)
		if err != nil {
			log.Error(err)
			return
		}
		log.Debugf("Write file from [%s] to [%s]", src_tmpl_path, dest_tmpl_path)
	}

	err = DirectoryFromTemplate(ISSUES_PATH, path.Join(TEMPLATES_DIR, "ISSUE_TEMPLATES"))
	if err != nil {
		log.Errorf("ISSUES: %v\n", err)
		return
	}

	return
}

func DirectoryFromTemplate(directory, template string) error {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}

	tmpl_files, err := tplDir.ReadDir(template)
	if err != nil {
		return err
	}

	for _, tmpl := range tmpl_files {
		src_tmpl_path := path.Join(template, tmpl.Name())

		fileContent, err := tplDir.ReadFile(src_tmpl_path)
		if err != nil {
			return err
		}
		dest_tmpl_path := path.Join(directory, tmpl.Name())
		err = os.WriteFile(dest_tmpl_path, fileContent, 0666)
		if err != nil {
			return err
		}
		log.Debugf("Write file %s", dest_tmpl_path)
	}
	return nil
}
