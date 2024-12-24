package main

import (
	"log"

	"github.com/go-l2-tasks/develop/dev11/pkg/server"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath("conf/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}
	err := server.StartServer(viper.GetString("port"))
	if err != nil {
		log.Fatal(err)
	}
}
