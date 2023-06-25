package controller

import (
	"github.com/gin-gonic/gin"

	"masmaint/core/auth"
)


func SetRouter(r *gin.Engine) {

	rm := r.Group("/mastertables", auth.NoopAuthMiddleware())
	{
		rm.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", gin.H{})
		})
	}
}