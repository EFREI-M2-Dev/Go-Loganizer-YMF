package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "loganizer",
	Short: "Loganizer est un outil (CLI) permettant d'annaliser des logs de diverses sources.",
	Long: "Loganizer est un outil (CLI) permettant d'analyser des logs de diverses sources. L'objectif est de pouvoir centraliser l'analyse de multiples logs en parallèle et d'en extraire des informations clés, tout en gérant les erreurs potentielles de manière robuste.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}
}