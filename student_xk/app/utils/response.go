package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatusCode int, code int, msg string, data interface{}) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, 200, "OK", data)
}

func ResponseInternalError(c *gin.Context) {
	Response(c, 200, 400, "Internal server error", nil)
}

func ResponseError(c *gin.Context, code int, msg string) {
	Response(c, http.StatusInternalServerError, code, msg, nil)
}
