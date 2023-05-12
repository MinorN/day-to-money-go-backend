package router

import (
	"mangosteen/config"
	"mangosteen/internal/controller"
	"mangosteen/internal/database"

	"mangosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	// 测试使用
	config.LoadViperConfig()
	r := gin.Default()
	docs.SwaggerInfo.Version = "1.0"

	database.Connect()

	r.GET("/ping", controller.Ping)

	r.POST("/api/v1/validation_codes", controller.CreateValidationCode)
	r.POST("/api/v1/session", controller.CreateSession)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
