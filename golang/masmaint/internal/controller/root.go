package controller

import (
	"github.com/gin-gonic/gin"
	
	"masmaint/internal/core/jwt"
)


type rootController struct {}


func newRootController() *rootController {
	return &rootController{}
}


//GET /
func (ctr *rootController) indexPage(c *gin.Context) {
	username := jwt.GetUserName(c)

	c.HTML(200, "index.html", gin.H{
		"username": username,
	})
}