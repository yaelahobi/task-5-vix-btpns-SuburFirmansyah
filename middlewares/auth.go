package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"task-5-vix-btpns-SuburFirmansyah/helpers"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := helpers.TokenValidation(c.Request, db)
		if status != http.StatusOK {
			if status == http.StatusUnauthorized {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return
			}
			return
		}
		c.Next()
	}
}
