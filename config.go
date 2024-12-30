package main

import (
	"encoding/json"
	"io"
	"os"
)

type config map[string]string

func newConfig(filePath string) (config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	context, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := config{}

	err = json.Unmarshal(context, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c config) getCommand(key string) (string, bool) {
	value, ok := c[key]
	return value, ok
}
