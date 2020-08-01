package controllers

import (
	"dl-admin-go/models"
	"dl-admin-go/services"
	"dl-admin-go/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(c *gin.Context) {

}

// @summary 用户登陆
// @Accept  json
// @Produce json
// @Param   user body models.User true "username&pwd"
// @Success 200 {string} string	"token"
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	if token, success := services.Login(user); success {
		c.String(http.StatusOK, token)
	} else {
		utils.NewError(c, http.StatusBadRequest, "username or password invalid")
	}
}

// @summary 获取当前的用户信息
// @Accept  json
// @Produce json
// @Success 200 {object} models.CustomClaims
// @Security ApiKeyAuth
// @Router /user [get]
func GetCurrentUser(c *gin.Context) {
	a, _ := c.Get("currentUser")

	if obj, success := a.(*models.CustomClaims); !success {
		utils.NewError(c, http.StatusNotFound, "user not found")
	} else {
		c.JSON(http.StatusOK, obj)
	}
}
