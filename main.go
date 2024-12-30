package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <file>", filepath.Base(os.Args[0]))
	}

	filePath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	fileBase := filepath.Base(filePath)
	fileExt := filepath.Ext(fileBase)
	if fileExt == "" {
		log.Fatal("Error: File must have an extension")
	}

	os.Setenv("FILE_PATH", filePath)
	os.Setenv("FILE_NAME", fileBase)
	os.Setenv("DIR", filepath.Dir(filePath))
	os.Setenv("FILE_NAME_NO_EXT", getFileNameNotExt(fileBase))

	profile, err := newConfig("config.json")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	command, ok := profile.getCommand(fileExt)
	if !ok {
		log.Fatalf("Error: No command found for file extension %s", fileExt)
	}

	execCommand(command)
}

func getFileNameNotExt(filePath string) string {
	for i, v := range filePath {
		if v == '.' {
			return filePath[:i]
		}
	}
	return ""
}
