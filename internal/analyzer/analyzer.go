package analyzer

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/config"
	"github.com/EFREI-M2-Dev/Go-Loganizer-YMF/internal/utils"
)

type CheckResult struct {
	InputTarget config.InputTarget
	Status string
	Err error
}

func AnalyzeLogFile(target config.InputTarget) CheckResult{
	fmt.Printf("[%s] Démarrage de l'analyse du log : %s\n", target.Id, target.Path)

	if _, err := os.Stat(target.Path); err != nil {
		fmt.Printf("[%s] Erreur : impossible d'accéder au fichier (%v)\n", target.Id, err)
		return CheckResult{
			InputTarget: target,
			Status: "error",
			Err: fmt.Errorf("impossible d'accéder au fichier %s : %w",
				target.Path, err),
		}
	}

	sleepMs := utils.RandomRange(50, 200)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	if rand.Float64() < 0.1 {
		fmt.Printf("[%s] Erreur : parsing du log échoué\n", target.Id)
		return CheckResult{
			InputTarget: target,
			Status: "error",
			Err: fmt.Errorf("parsing du log échoué pour le fichier %s", target.Path),
		}
	}

	fmt.Printf("[%s] Analyse terminée avec succès !\n", target.Id)

	return CheckResult{
		InputTarget: target,
		Status: "success",
		Err: nil,
	}

}