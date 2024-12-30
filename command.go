package main

import (
	"fmt"
	"os/exec"
)

func formatCommand(command string) string {
	newCommand := []rune(command)
	for i, v := range command {
		if v == '{' || v == '}' {
			newCommand[i] = '%'
		}
	}
	return string(newCommand)
}

func execCommand(command string) {
	formattedCommand := formatCommand(command)

	cmd := exec.Command("cmd", "/c", formattedCommand)
	output, _ := cmd.CombinedOutput()

	fmt.Print(string(output))
}
