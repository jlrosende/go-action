/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"io"
	"os"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	types "gitlab.com/jlrosende/go-action/types"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile    string
	config     = &types.Config{}
	log_level  string
	log_format string
	rootCmd    = &cobra.Command{
		Use:     "sisu",
		Short:   "Tool to deploy and test Sisu Functions",
		Long:    `.`,
		Version: "v1.0.0",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := setUpLogs(os.Stdout, log_level); err != nil {
				return err
			}
			return nil
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is sisu.{yml,yaml})")
	rootCmd.PersistentFlags().StringVarP(&log_level, "log-level", "l", log.WarnLevel.String(), "Log level (trace, debug, info, warn, error, fatal, panic")
	rootCmd.PersistentFlags().StringVar(&log_format, "log-format", "", "Log format (logfmt, json, text)")

}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("sisu")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("unable to unmarshall the config %v", err)
	}

	log.Warnf("%+v", *config)

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		for i, e := range err.(validator.ValidationErrors) {
			log.Errorf("Missing required attributes %d %v\n", i, e)
		}
		log.Fatal()
	}

	// for _, v := range config.Environments {
	// 	if err := validate.Struct(v); err != nil {
	// 		log.Fatalf("Missing required attributes %v\n", err)
	// 	}
	// }

	if c, err := yaml.Marshal(*config); err != nil {
		log.Fatalf("Missing required attributes %v\n", err)
	} else {
		log.Fatalf("\n%+v", string(c))
	}
}

func setUpLogs(out io.Writer, level string) error {
	log.SetOutput(out)
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	switch log_format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}

	return nil
}
