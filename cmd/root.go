/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jlrosende/go-action/cmd/create"
	conf "github.com/jlrosende/go-action/config"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile    string
	config     conf.Config
	log_level  string
	log_format string
	rootCmd    = &cobra.Command{
		Use:     "sisu",
		Short:   "Tool to interact with Sisu",
		Long:    `.`,
		Version: "v1.0.0",
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
	rootCmd.PersistentFlags().StringVarP(&log_level, "log-level", "l", log.InfoLevel.String(), "Log level (trace, debug, info, warn, error, fatal, panic")

	rootCmd.PersistentFlags().StringVar(&log_format, "log-format", "", "Log format (logfmt, json, text)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is sisu.{yml,yaml})")

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(matrixCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(updateCmd)

}

func initConfig() {

	if err := setUpLogs(os.Stdout, log_level); err != nil {
		log.Fatal(err)
	}

	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	viper.AddConfigPath(".")
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName("sisu")
	// }

	// viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err == nil {
	// 	log.Debugf("Using config file: %s", viper.ConfigFileUsed())
	// }

	// if err := viper.Unmarshal(config); err != nil {
	// 	log.Fatalf("unable to unmarshall the config %v", err)
	// }

	// validate := validator.New()
	// if err := validate.Struct(config); err != nil && viper.ConfigFileUsed() != "" {
	// 	log.Fatalf("Missing required attributes %v\n", err)
	// }

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
