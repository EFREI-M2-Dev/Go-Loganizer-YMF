package cmd

import (
	"fmt"

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

		fmt.Printf("Analyse des logs à partir du fichier : %s\n", configFilePath)
	},
}

func init() {
	analyzeCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "Chemin du fichier de configuration à utiliser")
	rootCmd.AddCommand(analyzeCmd)
}
