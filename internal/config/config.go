package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InputTarget struct {
	Id 	  string `json:"id"`
	Path  string `json:"path"`
	Type  string `json:"type"`
}

func LoadTargetsFromFile(filePath string) ([]InputTarget, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	var targets []InputTarget
	if err := json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("failed to unmarshall from file %s: %w", filePath, err)
	}
	return targets, nil
}

func AuthorizedLogFileTypes() []string {
	return []string{"nginx-access", "nginx-error", "custom-app", "generic", "mysql-error"}
}