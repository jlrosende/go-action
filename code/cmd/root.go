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
	initMod "github.com/jlrosende/go-action/cmd/init"
	"github.com/jlrosende/go-action/cmd/matrix"
	"github.com/jlrosende/go-action/cmd/test"
	"github.com/jlrosende/go-action/cmd/update"
	conf "github.com/jlrosende/go-action/config"
)

// rootCmd represents the base command when called without any subcommands
var (
	// cfgFile string
	// config     conf.Config
	log_level  string
	log_format string
	rootCmd    = &cobra.Command{
		Use:     "sisu",
		Short:   "Tool to interact with Sisu",
		Long:    `.`,
		Version: conf.Version,
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

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(matrix.MatrixCmd)
	rootCmd.AddCommand(initMod.InitCmd)
	rootCmd.AddCommand(test.TestCmd)
	rootCmd.AddCommand(update.UpdateCmd)

	if os.Getenv("GITHUB_OUTPUT") == "" {
		tmp, err := os.CreateTemp(os.TempDir(), "sisu-*.out")
		if err != nil {
			log.Fatalf("ERROR: %s", err)
			return
		}
		err = os.Setenv("GITHUB_OUTPUT", tmp.Name())
		if err != nil {
			log.Fatalf("ERROR: %s", err)
			return
		}
	}
}

func initConfig() {

	if err := setUpLogs(os.Stdout, log_level); err != nil {
		log.Fatal(err)
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
