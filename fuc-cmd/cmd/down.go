/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"fuc/fuc-cmd/cmd/utils"
	"os"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Destroy infrastructure for a specified environment",
	Long: `The down command destroys all infrastructure resources
created for a given environment.

It removes containers, networks, and any associated resources
managed by Terraform.

Examples:

  fuc down dev
  fuc down prod

This command will safely tear down all resources associated
with the selected environment.`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}

		environment := args[0]
		fmt.Println("Destroying infrastructure for environment:", environment)

		execDirectory := utils.GetCoreExecDir()

		var backendConfig string
		var tfvars string

		if environment == "dev" {
			backendConfig = "../../environments/dev/02_core.backend.config"
			tfvars = "../../environments/dev/dev.tfvars"
		} else if environment == "prod" {
			backendConfig = "../../environments/prod/02_core.backend.config"
			tfvars = "../../environments/prod/prod.tfvars"
		} else {
			fmt.Println("Invalid environment. Use 'dev' or 'prod'")
			os.Exit(1)
		}

		utils.RunTerraform(execDirectory, "init", "-reconfigure", "-backend-config="+backendConfig)
		utils.RunTerraform(execDirectory, "destroy", "-var-file="+tfvars, "--auto-approve")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
