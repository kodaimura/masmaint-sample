package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"masmaint/config"
	"masmaint/internal/core/jwt"
	"masmaint/internal/middleware"

	"masmaint/internal/module/employee"
	"masmaint/internal/module/department"
)

/*
 Routing for "/" 
*/
func SetWebRouter(r *gin.RouterGroup) {
	employeeController := employee.NewController()
	departmentController := department.NewController()

	r.GET("/login", func(c *gin.Context) { c.HTML(200, "login.html", gin.H{}) })

	auth := r.Group("", middleware.JwtAuthMiddleware())
	{
		auth.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", gin.H{}) })
		auth.GET("/employee", employeeController.GetPage)
		auth.GET("/department", departmentController.GetPage)
	}
}


func SetApiRouter(r *gin.RouterGroup) {
	employeeController := employee.NewController()
	departmentController := department.NewController()

	//カスタム推奨
	r.POST("/login", func(c *gin.Context) { 
		var body map[string]string
		c.ShouldBindJSON(&body)
		name := body["username"]
		pass := body["password"]
		fmt.Println(name)

		fmt.Println(pass)

		cf := config.GetConfig()
		if name == cf.AuthUser && pass == cf.AuthPass {
			cc := jwt.CustomClaims{ AccountId: 1, AccountName: name}
			jwt.SetTokenToCookie(c, jwt.NewPayload(cc))
		} else {
			c.JSON(401, gin.H{"error": "ユーザ名またはパスワードが異なります。"})
		}
	})

	auth := r.Group("", middleware.JwtAuthApiMiddleware())
	{
		auth.GET("/employee", employeeController.Get)
		auth.POST("/employee", employeeController.Post)
		auth.PUT("/employee", employeeController.Put)
		auth.DELETE("/employee", employeeController.Delete)

		auth.GET("/department", departmentController.Get)
		auth.POST("/department", departmentController.Post)
		auth.PUT("/department", departmentController.Put)
		auth.DELETE("/department", departmentController.Delete)
	}
}