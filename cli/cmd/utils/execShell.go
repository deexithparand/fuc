package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecShell(containerID string) {
	command := exec.Command("docker", "exec", "-it", containerID, "sh")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		fmt.Println("Error running docker:", err)
		os.Exit(1)
	}
}
