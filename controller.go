package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func runCheck(command string) int {
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}
	if err != nil {
		return 1
	}
	return 0
}
