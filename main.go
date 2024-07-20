package main

import (
	"github.com/lucassf2k/gopportunities-back/config"
	"github.com/lucassf2k/gopportunities-back/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
