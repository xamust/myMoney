package main

import (
	"github.com/xamust/myMoney.git/config"
	"github.com/xamust/myMoney.git/internal/app"
	"log"
)

func main() {
	// Get configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// Run App
	app.Run(cfg)
}
