/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "you're at the right place to setup linux on top of any os",
	Long:  `linuxtf says how fucking easy is it to setup linux on top on any OS using a type 2 virtualization concept such as docker container engine"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("linuxtf says hello")
		executeCmd := exec.Command("terraform", "state", "list")
		executeCmd.Dir = "../terraform/infra/02_core"
		bytes, err := executeCmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println("output : ")
		fmt.Println(string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
