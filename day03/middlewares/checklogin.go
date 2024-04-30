package middlewares

import (
	"SoftwareGoDay3/routes/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user jwt.User
		err := c.ShouldBindBodyWith(&user, binding.JSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}