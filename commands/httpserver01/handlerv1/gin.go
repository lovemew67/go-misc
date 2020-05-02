package handlerv1

import (
	"github.com/gin-gonic/gin"
)

func AddHTTPEndpoint(router *gin.RouterGroup) {
	router.Any("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello from": "v1 handlers",
			"server":     "http server 01",
		})
	})
}
