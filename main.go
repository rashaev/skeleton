package main

import (
	"github.com/rashaev/skeleton/internal/config"
	"github.com/rashaev/skeleton/internal/logger"
)

func main() {
	cfg := config.InitConfig()
	log := logger.InitLogger(cfg.Logging.Logfile, cfg.Logging.Severity)
	log.Info("Running service  ", cfg.App.Host, ":", cfg.App.Port)

}
