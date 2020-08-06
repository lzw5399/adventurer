package router

import (
	"adventurer/controller"
	"adventurer/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// default allow all origins
	r.Use(cors.Default())

	// swagger
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	account := r.Group("/account")
	{
		account.POST("register", controller.Register)
		account.POST("login", controller.Login)
		account.GET("user", middleware.Auth(), controller.GetCurrentUser)
	}

	user := r.Group("/user", middleware.Auth())
	{
		user.GET("list", controller.ListUsers)
	}

	return r
}
