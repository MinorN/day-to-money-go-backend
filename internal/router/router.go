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

func loadControllers() []controller.Controller {
	return []controller.Controller{
		&controller.SessionController{},
		&controller.ValidationCodeController{},
	}
}

func New() *gin.Engine {
	// 测试使用
	config.LoadViperConfig()
	r := gin.Default()
	docs.SwaggerInfo.Version = "1.0"

	database.Connect()

	r.GET("/ping", controller.Ping)

	api := r.Group("/api")
	cs := loadControllers()
	for _, c := range cs {
		c.RegisterRoutes(api)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
