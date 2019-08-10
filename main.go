package main

import (
	//"fmt"
	"github.com/rashaev/skeleton/internal/config"
	"github.com/rashaev/skeleton/internal/logger"
)

func main(){
	cfg := config.InitConfig()
	InitLogger(cfg.Logging.Logfile, cfg.Logging.Severity)
	//fmt.Println(cfg.App.Port)
	//log.Info("Running service")
}