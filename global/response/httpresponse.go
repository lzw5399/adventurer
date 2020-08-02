/**
 * @Author: lzw5399
 * @Date: 2020/7/31 23:21
 * @Desc: 格式化错误响应
 */
package response

import (
	"adventurer/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HTTPResponse struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context) {
	result(c, true, nil, "success")
}

func OkWithMessage(c *gin.Context, message string) {
	result(c, true, nil, message)
}

func OkWithData(c *gin.Context, data interface{}) {
	result(c, true, data, "操作成功")
}

func OkWithDetailed(c *gin.Context, data interface{}, message string) {
	result(c, true, data, message)
}

func Failed(c *gin.Context, status int) {
	FailWithMessage(c, status, "操作失败")
}

func FailWithMessage(c *gin.Context, status int, err interface{}) {
	var msg interface{}

	switch err.(type) {
	case validator.ValidationErrors:
		msg = err.(validator.ValidationErrors).Translate(global.DL_TRANSLOTOR)

	case error:
		msg = err.(error).Error()

	case string:
		msg = err.(string)

	default:
		msg = ""
	}

	// TODO: 这一块之后需要根据status code来做判断
	// 500: 的话不能直接返回，而是记日志再返回通用的错误信息
	switch status {
	// 400
	case http.StatusBadRequest:

	// 401
	case http.StatusUnauthorized:

	// 404
	case http.StatusNotFound:

	// 500
	case http.StatusInternalServerError:
		msg = "server internal error, please contact the maintainer"
	}

	resultWithStatus(c, status, false, nil, msg)
	c.Abort()
}

func result(c *gin.Context, success bool, data interface{}, msg string) {
	c.JSON(http.StatusOK, HTTPResponse{
		success,
		data,
		msg,
	})
}

func resultWithStatus(c *gin.Context, statusCode int, success bool, data interface{}, msg interface{}) {
	c.JSON(statusCode, HTTPResponse{
		success,
		msg,
		data,
	})
}
