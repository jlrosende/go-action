package update

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	conf "github.com/jlrosende/go-action/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	cfgFile   string
	UpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "",
		Long:  ``,
		Run:   update,
	}
)

type Workflow struct {
	Env Env `yaml:"env"`
}

type Env struct {
	WORKFLOW_VERSION string `yaml:"WORKFLOW_VERSION"`
	UPDATE_WORKFLOW  bool   `yaml:"UPDATE_WORKFLOW"`
}

func init() {

}

func update(cmd *cobra.Command, args []string) {

	log.Trace("cfgFile", cfgFile)
	config, err := conf.LoadConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Tracef("%+v", config)

	fmt.Println("Update")
	archivos, err := ioutil.ReadDir(".github/workflows")
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		if strings.HasPrefix(archivo.Name(), "sisu-") {
			log.Debug("Name: ", archivo.Name())

			body, err := ioutil.ReadFile(".github/workflows/" + archivo.Name())
			if err != nil {
				return
			}
			var w Workflow
			err = yaml.Unmarshal(body, &w)
			if err != nil {
				log.Errorf("ERROR: %v\n", err)
				return
			}
			log.Debugf("%+v", w)

			err = filepath.Walk(".github/temp/templates/workflows", func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {

					if info.Name() == archivo.Name() {
						// log.Trace("Path: ", path, "Info: ", info)
						if w.Env.WORKFLOW_VERSION != conf.Version && w.Env.UPDATE_WORKFLOW {
							log.Trace("Path: ", path, "Info: ", info)
							body, err := ioutil.ReadFile(path)
							if err != nil {
								return err
							}
							err = os.WriteFile(".github/workflows/"+archivo.Name(), body, 0666)
							if err != nil {
								return err
							}

						}
					}
				}

				return nil
			})
			if err != nil {
				log.Errorf("ERROR: %v\n", err)
				return
			}

		}
	}

	return
}
