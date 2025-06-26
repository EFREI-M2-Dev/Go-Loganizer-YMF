package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/analyzer"
	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/config"
	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/utils"
	"github.com/spf13/cobra"
)

var (
	configFilePath string
	outputFilePath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les logs fournis en entrée",
	Long: `Cette commande permet d'analyser les logs fournis en entrée.
Elle peut traiter des logs provenant de différentes sources et formats, en extrayant des informations pertinentes et en gérant les erreurs potentielles.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configFilePath == "" {
			fmt.Println("Veuillez spécifier un fichier d'entrée avec l'option (--config)")
			return
		}

		targets, err := config.LoadTargetsFromFile(configFilePath)
		if err != nil {
			fmt.Printf("Erreur lors du chargement du fichier de configuration : %v\n", err)
			return
		}

		var wg sync.WaitGroup
		var results []utils.CheckResult

		resultsChan := make(chan analyzer.CheckResult, len(targets))

		for _, target := range targets {
			wg.Add(1)
			go func(t config.InputTarget) {
				defer wg.Done()
				result := analyzer.AnalyzeLogFile(t)
				resultsChan <- result
			}(target)
		}

		wg.Wait()
		close(resultsChan)

		for result := range resultsChan {
			utilsResult := utils.CheckResult{
				InputTarget: result.InputTarget,
				Status:      result.Status,
				Timestamp:   time.Now(),
			}
			if result.Err != nil {
				utilsResult.Error = result.Err.Error()
			}
			results = append(results, utilsResult)
		}

		fmt.Println("Analyse de tous les logs terminée.")

		if outputFilePath != "" {
			if err := utils.ExportResultsToJSON(results, outputFilePath); err != nil {
				fmt.Printf("Erreur lors de l'export JSON : %v\n", err)
			}
		} else {
			fmt.Println("Aucun fichier de sortie spécifié. Utilisez --output pour exporter les résultats en JSON.")
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "Chemin du fichier de configuration à utiliser")
	analyzeCmd.MarkFlagRequired("config")
	analyzeCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Chemin du fichier JSON de sortie pour les résultats")
}
