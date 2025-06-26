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
		return CheckResult{
			InputTarget: target,
			Status: "error",
			Err: &UnreachableFileError{Path: target.Path, Err: err, Type: target.Type},
		}
	}

	if !utils.Contains(config.AuthorizedLogFileTypes(), target.Type) {
		return CheckResult{
			InputTarget: target,
			Status: "error",
			Err: &UnsupportedFileTypeError{Path: target.Path, Type: target.Type},
		}
	}

	sleepMs := utils.RandomRange(50, 200)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	if rand.Float64() < 0.1 {
		return CheckResult{
			InputTarget: target,
			Status: "error",
			Err: &ParsingError{Path: target.Path, Err: fmt.Errorf("parsing error"), Type: target.Type},
		}
	}

	fmt.Printf("[%s] Analyse terminée avec succès !\n", target.Id)

	return CheckResult{
		InputTarget: target,
		Status: "success",
		Err: nil,
	}

}