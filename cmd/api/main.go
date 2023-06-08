package main

import (
	"ShortURL/config"
	"ShortURL/internal/api"
	"ShortURL/internal/logging"
)

func main() {
	// Get logger for to display and save all api behavior
	log := logging.GetLogger()

	// Read all configurations from the yaml config file using the viper module
	err := config.GetConfigs()
	if err != nil {
		return
	}

	// Create and initialize API structure with all dependencies
	log.Info("api initialization ...")
	API := api.NewApi(&log)
	err = API.Init()
	if err != nil {
		return
	}

	log.Info("api initialized")
	//Start our API
	log.Info("api Running...")
	err = API.Start()
	if err != nil {
		return
	}
}
