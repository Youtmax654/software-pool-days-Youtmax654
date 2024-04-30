package routes

import (
	"SoftwareGoDay3/middlewares"
	"SoftwareGoDay3/routes/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyLogin(r *gin.Engine) {
	r.POST("/jwt/login", middlewares.CheckLogin(), login)
}

func login(c *gin.Context) {
	user := c.MustGet("user").(jwt.User)

	if !jwt.IsRegistered(user.Email) {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	if !jwt.IsPasswordCorrect(user.Email, user.Password) {
		c.String(http.StatusNotFound, "Wrong password")
		return
	}

	token, err := jwt.GenerateToken(user.Email, user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.JSON(http.StatusOK, jwt.AuthResponse{
		AccessToken: token,
		User:        user,
		Message:     "Successful login",
	})
}