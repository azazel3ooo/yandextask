package checkers

import (
	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/staticcheck"
	"strings"
)

// GetAllSA - формирует массив из всех анализаторов с префиксом SA
func GetAllSA() []*analysis.Analyzer {
	var sa []*analysis.Analyzer
	for _, v := range staticcheck.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, "SA") {
			sa = append(sa, v.Analyzer)
		}
	}

	return sa
}

// GetOtherStaticChecks - формирует массив из различных анализаторов пакета staticcheck
func GetOtherStaticChecks() []*analysis.Analyzer {
	checks := map[string]bool{
		"ST1020": true,
	}
	var res []*analysis.Analyzer
	for _, v := range staticcheck.Analyzers {
		if checks[v.Analyzer.Name] {
			res = append(res, v.Analyzer)
		}
	}

	return res
}
