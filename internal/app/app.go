package app

import (
	"github.com/xamust/myMoney.git/config"
	"github.com/xamust/myMoney.git/pkg/logger"
)

func Run(cfg *config.Config) {
	// Init logger
	logger.NewLogger(logger.WithLevel(cfg.Log.Level))

	// Init repository

}
