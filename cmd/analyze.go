package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/config"
	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/utils"
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

		for _, target := range targets {
			wg.Add(1)
			go func(t config.InputTarget) {
				defer wg.Done()
				fmt.Printf("[%s] Démarrage de l'analyse du log : %s\n", t.Id, t.Path)

				if _, err := os.Stat(t.Path); err != nil {
					fmt.Printf("[%s] Erreur : impossible d'accéder au fichier (%v)\n", t.Id, err)
					return
				}

				sleepMs := utils.RandomRange(50, 200)
				time.Sleep(time.Duration(sleepMs) * time.Millisecond)

				if rand.Float64() < 0.1 {
					fmt.Printf("[%s] Erreur : parsing du log échoué\n", t.Id)
					return
				}

				fmt.Printf("[%s] Analyse terminée avec succès !\n", t.Id)
			}(target)
		}
		wg.Wait()
		fmt.Println("Analyse de tous les logs terminée.")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "Chemin du fichier de configuration à utiliser")
	analyzeCmd.MarkFlagRequired("config")
}
