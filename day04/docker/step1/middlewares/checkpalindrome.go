package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CheckPalindrome() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputs []string
		err := c.ShouldBindBodyWith(&inputs, binding.JSON); 
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			c.Abort()
			return
		}

		c.Set("inputs", inputs)

		c.Next()
	}
}