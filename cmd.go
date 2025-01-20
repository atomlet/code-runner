package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func registerEnvironments(envVars map[string]string) {
	for key, value := range envVars {
		os.Setenv(key, value)
	}
}

func parseConfig() (map[string]string, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config map[string]string
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

func formatCommand(command string) string {
	runes := []rune(command)
	for i, char := range runes {
		if char == '{' || char == '}' {
			runes[i] = '%'
		}
	}
	return string(runes)
}

func extractFileName(path string) string {
	base := filepath.Base(path)
	if ext := filepath.Ext(base); ext != "" {
		return base[:len(base)-len(ext)]
	}
	return base
}

func executeCommand(command string) error {
	cmd := exec.Command("cmd", "/k", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func fatalln(args ...any) {
	fmt.Println(args...)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		fatalln("Usage: go run main.go <file_path>")
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		fatalln("Failed to resolve file path:", err)
	}

	config, err := parseConfig()
	if err != nil {
		fatalln("Error reading configuration:", err)
	}

	environments := map[string]string{
		"FILE_PATH": path,
		"FILE_NAME": extractFileName(path),
		"DIR":       filepath.Dir(path),
	}

	ext := filepath.Ext(path)
	command, exists := config[ext]
	if !exists {
		fatalln("No command found for file type:", ext)
	}

	registerEnvironments(environments)
	command = formatCommand(command)

	if err := executeCommand(command); err != nil {
		fatalln("Command execution failed:", err)
	}
}
