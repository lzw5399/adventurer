package controllers

import (
	"dl-admin-go/models"
	"dl-admin-go/services"
	"dl-admin-go/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

// @summary 用户登陆
// @Accept  json
// @Produce json
// @Param   user body models.User true "username&pwd"
// @Success 200 {string} string	"token"
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "model binding failed",
		})
	}

	if token, success := services.Login(user); success {
		c.String(http.StatusOK, token)
	} else {
		utils.NewError(c, http.StatusUnauthorized, "login failed")
	}
}

// @summary 获取当前的用户信息
// @Accept  json
// @Produce json
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /user [get]
func GetCurrentUser(c *gin.Context) {
	a, _ := c.Get("currentUser")

	if obj, success := a.(*models.CustomClaims); !success {
		c.PureJSON(http.StatusOK, gin.H{
			"msg": "认证成功但是无法获取当前用户",
		})
	} else {
		c.JSON(http.StatusOK, obj)
	}
}
