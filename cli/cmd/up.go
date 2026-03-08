/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"loc/cli/cmd/utils"
	"os"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Provision and start infrastructure for a specified environment",
	Long: `The up command provisions and starts the infrastructure defined 
for a given environment.

It pulls required container images, creates containers, and starts 
the services defined in the configuration.

Examples:

  loc up dev
  loc up prod

This command ensures that all required resources are created and 
running for the selected environment.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Simulate the below
		// if environment == "dev" {
		// 	// development

		// 	// terraform init --backend-config="../../environments/dev/02_core.backend.config"
		// 	// terraform plan -var-file="../../environments/dev/dev.tfvars"
		// 	// terraform apply -var-file="../../environments/dev/dev.tfvars" --auto-approve
		// } else {
		// 	// production

		// 	// terraform init --backend-config="../../environments/prod/02_core.backend.config"
		// 	// terraform plan -var-file="../../environments/prod/prod.tfvars"
		// 	// terraform apply -var-file="../../environments/prod/prod.tfvars" --auto-approve
		// }

		if len(args) == 0 {
			cmd.Help()
			return
		}

		environment := args[0]
		fmt.Println("Running infrastructure for environment:", environment)

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
		utils.RunTerraform(execDirectory, "plan", "-var-file="+tfvars)
		utils.RunTerraform(execDirectory, "apply", "-var-file="+tfvars, "--auto-approve")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
