package main

import (
	"log"
	"mangosteen/cmd"

	"github.com/spf13/viper"
)

// @title           山竹记账 API
// @description     这是本项目的所有 API 的文档

// @contact.name   小N
// @contact.url    https://MinorN.com
// @contact.email  magicminorn@gmail.com

// @license.name  MIT
// @license.url

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth(JWT)

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	cmd.Run()
}
