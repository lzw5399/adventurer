package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/appleboy/gin-jwt/v2"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// default allow all origins
	r.Use(cors.Default())

	r.Use(jwt.)

	// routers
	authGroup := r.Group("/")
	anonymousGroup := r.Group("/")

	{
		authGroup.GET("a", Emm)
	}

	anonymousGroup.GET("b", Emm2)

	return r
}

func Emm3(c *gin.Context) {
	c.String(http.StatusOK, "Emm3")
}

func Emm(c *gin.Context) {
	c.String(http.StatusOK, "Emm")
}

func Emm2(c *gin.Context) {
	c.String(http.StatusOK, "Emm2")
}
