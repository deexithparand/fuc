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

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open an interactive shell inside a running container",
	Long: `The shell command allows you to open an interactive shell
session inside one of the running containers managed by FUC.

It will display a list of active containers and prompt you
to select the container you want to access. Once selected,
an interactive shell session will be started inside that container.

This is useful for debugging, inspecting running services,
or executing commands directly within the container environment.

Examples:

  fuc shell
  fuc shell dev [dev-container-name]
  fuc shell prod [prod-container-name]

In interactive mode, you will be prompted to choose the container
you want to access from the available running containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		environment := args[0]

		if (len(args) == 1) && environment == "dev" {
			fmt.Println("Choose any of the dev containers to log into : ")
			utils.ListContainerNames(environment)
			return
		} else if (len(args) == 1) && environment == "prod" {
			fmt.Println("Choose any of the prod containers to log into : ")
			utils.ListContainerNames(environment)
			return
		} else if environment == "dev" {
			containerID := args[1]
			// get into any of the dev containers
			utils.ExecShell(containerID)
			return
		} else if environment == "prod" {
			containerID := args[1]
			// get into any of the prod containers
			utils.ExecShell(containerID)
			return
		} else {
			fmt.Println("Invalid environment. Use 'dev' or 'prod'")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
