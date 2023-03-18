package main

import (
	"github.com/bernardoamim/go-jobs/config"
	"github.com/bernardoamim/go-jobs/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
