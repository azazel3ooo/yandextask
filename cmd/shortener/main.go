package main

import (
	"fmt"
	"github.com/azazel3ooo/yandextask/internal/service"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf(
		"Build version: %s\n"+
			"Build date: %s\n"+
			"Build commit: %s\n", buildVersion, buildDate, buildCommit)

	service.StartService()
}
