// Package commands defines and implements command-line commands and flags
// used by migrator. Commands and flags are implemented using Cobra.
package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/devopsdays/migrator/helpers/paths"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var myBanner = `

+-+-+-+-+-+-+
|  migrator |
+-+-+-+-+-+-+

`

// webdir is the path to the source files for the Hugo website
var webdir = paths.GetWebdir()

var cfgFile string

// Debug means should we run in debug mode. Duh.
var Debug bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "migrator",
	Short: "Migrate to the updated devopsdays theme",
	Long: `
Command-line utilities for migrating the devopsdays.org website
built with love by mattstratton in Go.

Complete documentation is available at https://github.com/devopsdays/migrator`,
	Run: func(cmd *cobra.Command, args []string) {
		mainPrompt()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(myBanner))

	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "enable debug mode")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".migrator") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
