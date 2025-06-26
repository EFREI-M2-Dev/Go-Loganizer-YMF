package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

func DisplayResultsSummary(results []CheckResult) {
	if len(results) == 0 {
		fmt.Println("Aucun résultat à afficher.")
		return
	}

	total := len(results)
	success := 0
	errors := 0
	var errorFiles []string

	for _, result := range results {
		if result.Status == "success" {
			success++
		} else {
			errors++
			errorFiles = append(errorFiles, fmt.Sprintf("  - %s (%s): %s",
				result.InputTarget.Id, result.InputTarget.Path, result.Error))
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("           RÉSUMÉ DE L'ANALYSE")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total de logs analysés   : %d\n", total)
	fmt.Printf("Succès                   : %d\n", success)
	fmt.Printf("Erreurs                  : %d\n", errors)

	if errors > 0 {
		fmt.Println("\nFichiers en erreur :")
		for _, errorFile := range errorFiles {
			fmt.Println(errorFile)
		}
	}

	successRate := float64(success) / float64(total) * 100
	fmt.Printf("\nTaux de réussite         : %.1f%%\n", successRate)
	fmt.Println(strings.Repeat("=", 50))
}
