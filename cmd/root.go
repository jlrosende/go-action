/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile    string
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

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is sisu.{yml,yaml})")
	rootCmd.PersistentFlags().StringVarP(&log_level, "log-level", "l", log.WarnLevel.String(), "Log level (trace, debug, info, warn, error, fatal, panic")
	rootCmd.PersistentFlags().StringVar(&log_format, "log-format", "", "Log format (logfmt, json, text)")

	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(testCmd)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
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
