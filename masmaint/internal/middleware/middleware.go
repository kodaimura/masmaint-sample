package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"masmaint/config"
	"masmaint/internal/core/jwt"
	"masmaint/internal/core/errs"
)


func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cf := config.GetConfig()

		user, pass, ok := c.Request.BasicAuth()
		if !ok || user != cf.BasicAuthUser || pass != cf.BasicAuthPass {
			c.Header("WWW-Authenticate", "Basic realm=Authorization Required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}


func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := jwt.Auth(c); err != nil {
			c.Redirect(303, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}


func JwtAuthApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := jwt.Auth(c); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}


func ApiResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			switch e := err.(type) {
			case errs.BadRequestError:
				c.JSON(http.StatusBadRequest, gin.H{"error": e.Error(), "field": e.Field, "type": e.Type})
			case errs.NotFoundError:
				c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
			case errs.UniqueConstraintError:
				c.JSON(http.StatusConflict, gin.H{"error": e.Error(), "column": e.Column})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
			}
		}
	}
}