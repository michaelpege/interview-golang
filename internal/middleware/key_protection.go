package middleware

import (
	"interview/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func KeyProtection() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("KEY")
		if key != "12345" {
			utils.SendResponse(c, http.StatusUnauthorized, "Not Authorized", nil)
			c.Abort()
		}
		c.Next()

	}
}
