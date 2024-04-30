package routes

import (
	"SoftwareGoDay3/middlewares"
	"SoftwareGoDay3/routes/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyRegister(r *gin.Engine) {
	r.POST("/jwt/register", middlewares.CheckRegister(), register)
}

func register(c *gin.Context) {
	user := c.MustGet("user").(jwt.User)

	if jwt.IsRegistered(user.Email) {
		c.String(http.StatusForbidden, "User already exists")
		return
	}

	response := jwt.Register(user)

	c.JSON(http.StatusCreated, response)
}