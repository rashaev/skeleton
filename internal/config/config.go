package config

import (
	"github.com/spf13/viper"
	"log"
)


type Config struct {
	Host string 
	Port int
}


func InitConfig() Config {
	var C Config{}
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