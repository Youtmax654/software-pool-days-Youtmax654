package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyHealth(r *gin.Engine) {
	r.GET("/health", health)
}

func health(c *gin.Context) {
	c.Status(http.StatusOK)
}