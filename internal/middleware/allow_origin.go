package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var CONST_ALLOW_ORIGIN map[string]bool

func init() {
	CONST_ALLOW_ORIGIN = make(map[string]bool, 0)
	CONST_ALLOW_ORIGIN["http://localhost:3000"] = true
}

//SetResponseHeader set response header for all requests
func SetResponseHeader() gin.HandlerFunc {
	return func(cg *gin.Context) {
		reqOrigin := cg.Request.Header.Get("Origin")
		if _, haveOrigin := CONST_ALLOW_ORIGIN[reqOrigin]; haveOrigin || strings.Contains(reqOrigin, "localhost.com") {
			cg.Header("Access-Control-Allow-Origin", reqOrigin)
			cg.Header("Connection", "keep-alive")
			cg.Header("Access-Control-Allow-Credentials", "true")
			cg.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			cg.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, SessionKey")
		}
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method //请求方法
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusNoContent, "Options Request!")
		}
		c.Next() //  处理请求
	}
}
