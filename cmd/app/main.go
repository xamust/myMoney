package main

import (
	"github.com/xamust/myMoney.git/config"
	"log"
)

func main() {
	// Get configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	_ = cfg
	// Run App
	// ...
}
