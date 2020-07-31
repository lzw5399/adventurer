/**
 * @Author: lzw5399
 * @Date: 2020/7/31 23:21
 * @Desc: 格式化错误响应
 */
package utils

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(ctx *gin.Context, status int, err interface{}) {
	er := HTTPError{
		Code: status,
	}

	switch err.(type) {
	case error:
		er.Message = err.(error).Error()
	case string:
		er.Message = err.(string)
	}

	ctx.JSON(status, er)
}
