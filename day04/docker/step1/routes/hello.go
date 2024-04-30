package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyHello(r *gin.Engine) {
	r.GET("/hello", hello)
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "world")
}