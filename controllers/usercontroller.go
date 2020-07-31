package controllers

import (
	"dl-admin-go/models"
	"dl-admin-go/services"
	"dl-admin-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// @summary 获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id path int true "userId"
// @Success 200 {string} string	"ok"
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
