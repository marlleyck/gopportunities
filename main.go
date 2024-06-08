package main

import (
	"github.com/marlleyck/gopportunities/config"
	"github.com/marlleyck/gopportunities/router"
)

var (
	logger *config.Logger
)


func main() {
	logger = config.GetLogger("main")

	// Initialiaze Configs
	err := config.Init()

	if (err != nil) {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}