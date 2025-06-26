package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/config"
)

type CheckResult struct {
	InputTarget config.InputTarget `json:"input_target"`
	Status      string             `json:"status"`
	Error       string             `json:"error,omitempty"`
	Timestamp   time.Time          `json:"timestamp"`
}

func RandomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func ExportResultsToJSON(results []CheckResult, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("failed to encode results to JSON: %w", err)
	}

	fmt.Printf("Résultats exportés vers : %s\n", outputPath)
	return nil
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
