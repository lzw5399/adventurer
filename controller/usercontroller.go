package controller

import (
	"dl-admin-go/global/response"
	"dl-admin-go/model"
	"dl-admin-go/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(c *gin.Context) {

}

// @summary 用户登陆
// @Accept  json
// @Produce json
// @Param   user body request.UserAddRequest true "username&pwd"
// @Success 200 {string} string	"token"
// @Router /login [post]
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage(c, http.StatusBadRequest, err)
		return
	}

	if token, success := service.Login(user); success {
		c.String(http.StatusOK, token)
	} else {
		response.FailWithMessage(c, http.StatusBadRequest, "username or password invalid")
	}
}

// @summary 获取当前的用户信息
// @Accept  json
// @Produce json
// @Success 200 {object} model.CustomClaims
// @Security ApiKeyAuth
// @Router /user [get]
func GetCurrentUser(c *gin.Context) {
	a, _ := c.Get("currentUser")

	if obj, success := a.(*model.CustomClaims); !success {
		response.FailWithMessage(c, http.StatusNotFound, "user not found")
	} else {
		c.JSON(http.StatusOK, obj)
	}
}
