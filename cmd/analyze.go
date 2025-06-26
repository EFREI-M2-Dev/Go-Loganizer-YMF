package cmd

import (
	"fmt"
	"sync"

	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/analyzer"
	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/config"
	"github.com/spf13/cobra"
)

var (
	configFilePath string
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

		fmt.Println("Analyse de tous les logs terminée.")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "Chemin du fichier de configuration à utiliser")
	analyzeCmd.MarkFlagRequired("config")
}
