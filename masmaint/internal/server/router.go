package server

import (
	"github.com/gin-gonic/gin"

	"masmaint/internal/middleware"
	"masmaint/internal/controller"
	"masmaint/config"
	"masmaint/internal/core/jwt"
)

/*
 Routing for "/" 
*/
func SetWebRouter(r *gin.RouterGroup) {
	ic := controller.NewIndexController()

	r.GET("/login", func(c *gin.Context) { c.HTML(200, "login.html", gin.H{}) })

	auth := r.Group("", middleware.JwtAuthMiddleware())
	{
		auth.GET("/", ic.IndexPage)
	}
}


/*
 Routing for "/api"
*/
func SetApiRouter(r *gin.RouterGroup) {
	ac := controller.NewAccountController()

	//カスタム推奨
	r.POST("/login", func(c *gin.Context) { 
		name := c.Param("username")
		pass := c.Param("password")

		cf := config.GetConfig()
		if name == cf.AuthUser && pass == cf.AuthPass {
			cc := jwt.CustomClaims{
				AccountId:   1,
				AccountName: name,
			}
			jwt.SetTokenToCookie(c, jwt.NewPayload(cc))
		} else {
			c.JSON(401, gin.H{"error": "ユーザ名またはパスワードが異なります。"})
		}
	})

	auth := r.Group("", middleware.JwtAuthApiMiddleware())
	{
		auth.GET("/accounts/me", ac.ApiGetOne)
		auth.PUT("/accounts/me/name", ac.ApiPutName)
		auth.PUT("/accounts/me/password", ac.ApiPutPassword)
		auth.DELETE("/accounts/me", ac.ApiDelete)
	}
}