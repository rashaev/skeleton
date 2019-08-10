package logger

import (
	"os"
	"github.com/sirupsen/logrus"
)



func InitLogger(logfile, severity string) {
	var log = logrus.New() 
	log.Formatter = new(logrus.TextFormatter)
	
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Out = os.Stdout
	}
	
	switch severity {
	case "info":
		log.Level = logrus.InfoLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "debug":
		log.Level = logrus.DebugLevel
	}
}

