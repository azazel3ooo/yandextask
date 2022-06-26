// package main осуществляет запуск кастомного чекера. В его состав входят все
// чекеры с префиксом SA, некоторые другие анализаторы пакета staticcheck
// публичный чекер https://github.com/gostaticanalysis/nilerr
// и анализатор, который проверяет отсутствие вызовы функции os.Exit() в функции main пакета main

// публичный чекер проверяет отсутвие участков подобного кода
// func f() error {
//	err := do()
//	if err != nil {
//		return nil // miss
//	}
// }
package main

import (
	"github.com/gostaticanalysis/nilerr"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"mycheck/checkers"
)

func main() {
	var checks []*analysis.Analyzer

	SA := checkers.GetAllSA()
	otherCheckers := checkers.GetOtherStaticChecks()
	checks = append(checks, SA...)                            // all SA* checks
	checks = append(checks, shadow.Analyzer, printf.Analyzer) // some passes checks
	checks = append(checks, otherCheckers...)                 // other static checkers
	checks = append(checks, checkers.ExitCheckAnalyzer)       // my checker
	checks = append(checks, nilerr.Analyzer)                  // public analyzer

	multichecker.Main(
		checks...,
	)
}
