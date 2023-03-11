package main

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lucasd-coder/business-service/config"
	"github.com/lucasd-coder/business-service/internal/app"
)

var cfg config.Config

func main() {
	profile := os.Getenv("GO_PROFILE")
	var path string

	switch profile {
	case "dev":
		path = "./config/config-dev.yml"
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
	config.ExportConfig(&cfg)

	app.Run(&cfg)
}
