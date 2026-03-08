package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func GetCoreExecDir() string {
	return "../terraform/infra/02_core"
}

func RunTerraform(dir string, args ...string) {

	command := exec.Command("terraform", args...)
	command.Dir = dir

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		fmt.Println("Error running terraform:", err)
		os.Exit(1)
	}
}
