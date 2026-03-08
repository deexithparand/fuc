/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loc",
	Short: "Local Ops Cloud — build and experiment with infrastructure locally",
	Long: `loc (Local Ops Cloud) is a lightweight CLI for building and experimenting
with infrastructure environments locally.

It provides a simple workflow to create, manage, and destroy environments
using infrastructure as code — without requiring a cloud account.

LOC is designed as a playground for learning and experimentation.
Spin up containers, simulate infrastructure setups, explore workflows,
and understand how infrastructure tools behave in a safe local environment.

Whether you're learning infrastructure concepts, testing ideas, or building
your own automation tools, LOC lets you experiment freely without worrying
about cloud costs.

Examples:
  loc up 	         Start the development environment
  loc down	         Tear down an environment
  loc shell          Access a running container
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var banner = `
██╗      ██████╗  ██████╗
██║     ██╔═══██╗██╔════╝
██║     ██║   ██║██║     
██║     ██║   ██║██║     
███████╗╚██████╔╝╚██████╗
╚══════╝ ╚═════╝  ╚═════╝

Local Ops Cloud
`

func init() {
	fmt.Println(banner)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.loc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
