package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	jwtHelper "picnshop/pkg/jwt"
)

func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey, os.Getenv("ENV"))
			if decodedClaims != nil && decodedClaims.IsAdmin {
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}

func AuthUserMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey, os.Getenv("ENV"))
			if decodedClaims != nil {
				c.Set("userId", decodedClaims.UserId)
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
