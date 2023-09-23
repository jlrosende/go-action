package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jlrosende/go-action/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var Version = "v2.0.0"

type Config struct {
	Version string                `yaml:"version" mapstructure:"version" validate:"required"`
	Env     map[string][]Function `yaml:"environments" mapstructure:"environments" validate:"required,dive,dive"`
}

func LoadConfig(cfgFile string) (*Config, error) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./cloud")
		viper.SetConfigType("yaml")
		viper.SetConfigName("sisu")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Debugf("Using config file: %s", viper.ConfigFileUsed())

	yamlFile, err := os.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(&config); err != nil && viper.ConfigFileUsed() != "" {
		return nil, err
	}

	return &config, nil
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) AddEnvironment(name string) {

}

func (c *Config) AddFunction(environment string, function Function) {
	c.Env[environment] = append(c.Env[environment], function)
}

func (c *Config) GetRuntimes() []string {
	runtimes := []string{}

	for _, funtions := range c.Env {
		for _, funtion := range funtions {
			runtimes = append(runtimes, funtion.Runtime)
		}
	}

	return utils.RemoveDuplicateStr(runtimes)
}

func (c *Config) GetLang() string {
	langs := []string{}

	for _, funtions := range c.Env {
		for _, funtion := range funtions {
			langs = append(langs, utils.ParseRuntime(funtion.Runtime))
		}
	}

	return utils.RemoveDuplicateStr(langs)[0]
}
