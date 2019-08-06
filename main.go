package main

import (
	"fmt"
	"github.com/rashaev/skeleton/internal/config"
)

func main(){
	cfg := config.InitConfig()
	fmt.Println(cfg.Host)
}