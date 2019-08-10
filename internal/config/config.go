package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"app"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Logging struct {
		Logfile  string `json:"logfile"`
		Severity string `json:"severity"`
	} `json:"logging"`
}

func InitConfig() Config {
	var C Config
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return C
}
