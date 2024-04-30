package routes

import (
	"SoftwareGoDay3/routes/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func applyMe(r *gin.Engine) {
	r.GET("/jwt/me", me)
}

func me(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	token := strings.Split(auth, "Bearer ")[1]

	user, err := jwt.GetUserFromToken(token)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	if !jwt.IsRegistered(user.Email) {
		c.String(http.StatusNotFound, "Unknown user")
		return
	}

	c.JSON(http.StatusOK, jwt.MeResponse{
		User:    user,
		Message: "User found",
	})
}