package server

import (
	"github.com/gin-gonic/gin"
	"masmaint/config"
	"masmaint/internal/core/jwt"
	"masmaint/internal/middleware"

	"masmaint/internal/module/customer"
	"masmaint/internal/module/product_category"
	"masmaint/internal/module/product"
	"masmaint/internal/module/supplier"
	"masmaint/internal/module/payment_method"
)

/*
 Routing for "/" 
*/
func SetWebRouter(r *gin.RouterGroup) {
	customerController := customer.NewController()
	product_categoryController := product_category.NewController()
	productController := product.NewController()
	supplierController := supplier.NewController()
	payment_methodController := payment_method.NewController()

	r.GET("/login", func(c *gin.Context) { c.HTML(200, "login.html", gin.H{}) })

	auth := r.Group("", middleware.JwtAuthMiddleware())
	{
		auth.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", gin.H{}) })
		auth.GET("/customer", customerController.GetPage)
		auth.GET("/product_category", product_categoryController.GetPage)
		auth.GET("/product", productController.GetPage)
		auth.GET("/supplier", supplierController.GetPage)
		auth.GET("/payment_method", payment_methodController.GetPage)
	}
}


func SetApiRouter(r *gin.RouterGroup) {
	r.Use(middleware.ApiResponseMiddleware())

	customerController := customer.NewController()
	product_categoryController := product_category.NewController()
	productController := product.NewController()
	supplierController := supplier.NewController()
	payment_methodController := payment_method.NewController()

	//カスタム推奨
	r.POST("/login", func(c *gin.Context) { 
		var body map[string]string
		c.ShouldBindJSON(&body)
		name := body["username"]
		pass := body["password"]

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
		auth.GET("/customer", customerController.Get)
		auth.POST("/customer", customerController.Post)
		auth.PUT("/customer", customerController.Put)
		auth.DELETE("/customer", customerController.Delete)

		auth.GET("/product_category", product_categoryController.Get)
		auth.POST("/product_category", product_categoryController.Post)
		auth.PUT("/product_category", product_categoryController.Put)
		auth.DELETE("/product_category", product_categoryController.Delete)

		auth.GET("/product", productController.Get)
		auth.POST("/product", productController.Post)
		auth.PUT("/product", productController.Put)
		auth.DELETE("/product", productController.Delete)

		auth.GET("/supplier", supplierController.Get)
		auth.POST("/supplier", supplierController.Post)
		auth.PUT("/supplier", supplierController.Put)
		auth.DELETE("/supplier", supplierController.Delete)

		auth.GET("/payment_method", payment_methodController.Get)
		auth.POST("/payment_method", payment_methodController.Post)
		auth.PUT("/payment_method", payment_methodController.Put)
		auth.DELETE("/payment_method", payment_methodController.Delete)
	}
}