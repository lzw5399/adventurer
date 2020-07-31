package routers

import (
	"dl-admin-go/controllers"
	"dl-admin-go/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// default allow all origins
	r.Use(cors.Default())

	anonymousGroup := r.Group("/")
	{
		anonymousGroup.POST("login", controllers.Login)
	}

	authGroup := r.Group("/", middlewares.Auth())
	{
		authGroup.GET("user", controllers.GetCurrentUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
