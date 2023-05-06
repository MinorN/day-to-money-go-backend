package router

import (
	"mangosteen/config"
	"mangosteen/internal/controller"

	"mangosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	// 测试使用
	config.LoadViperConfig()

	docs.SwaggerInfo.Version = "1.0"

	r := gin.Default()

	r.GET("/api/v1/ping", controller.Ping)

	r.POST("/api/v1/validation_codes", controller.CreateValidationCode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
