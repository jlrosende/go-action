package config

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Version string                `yaml:"version" mapstructure:"version" validate:"required"`
	Env     map[string][]Function `yaml:"environments" mapstructure:"environments" validate:"required,dive,dive"`
}

func LoadConfig(cfgFile string) (*Config, error) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("./cloud")
		viper.SetConfigType("yaml")
		viper.SetConfigName("sisu")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debugf("Using config file: %s", viper.ConfigFileUsed())
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		// log.Fatalf("unable to unmarshall the config %v", err)
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(&config); err != nil && viper.ConfigFileUsed() != "" {
		// log.Fatalf("Missing required attributes %v\n", err)
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
