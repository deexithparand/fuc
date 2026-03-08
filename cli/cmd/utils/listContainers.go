package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ListContainerNames(environment string) {
	command := exec.Command("docker", "ps")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		fmt.Println("Error running docker:", err)
		os.Exit(1)
	}
}
