package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadViperConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
