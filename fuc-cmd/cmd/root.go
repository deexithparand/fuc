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
	Use:   "fuc",
	Short: "Forge Your Cloud — spin up dev and prod environments instantly",
	Long: `FUC (Forge Your Cloud) is a developer-focused CLI that helps you
create, manage, and destroy infrastructure environments with minimal effort.

It provides a simple workflow to bootstrap containers, services,
and environments like dev and prod using infrastructure as code.

Whether you're testing locally or simulating production setups,
FUC makes environment orchestration fast, reproducible, and simple.

Examples:

  fuc init           Initialize a FUC project
  fuc up dev         Start the development environment
  fuc up prod        Start the production environment
  fuc down           Tear down running infrastructure
  fuc shell          Access a running container
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
███████╗██╗   ██╗ ██████╗
██╔════╝██║   ██║██╔════╝
█████╗  ██║   ██║██║     
██╔══╝  ██║   ██║██║     
██║     ╚██████╔╝╚██████╗
╚═╝      ╚═════╝  ╚═════╝

Forge Your Cloud
`

func init() {
	fmt.Println(banner)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fuc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
