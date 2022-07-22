package main

import (
	"avito-test/internal/app"
	"avito-test/internal/config"
	"avito-test/pkg/logging"
	"log"
)

func main() {
	log.Print("config initialization")
	cfg := config.GetConfig()

	log.Printf("logger initialized.")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
	a.Run()
}
