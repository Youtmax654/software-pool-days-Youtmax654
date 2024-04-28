package routes

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.Engine) {
	router.GET("/hello", hello)
	router.GET("/repeat-my-query", repeatMyQuery)
	router.GET("/repeat-my-param/:message", repeatMyParam)
	router.POST("/repeat-my-body", repeatMyBody)
	router.GET("/repeat-my-header", repeatMyHeader)
	router.GET("/repeat-my-cookie", repeatMyCookie)
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "world")
}

func repeatMyQuery(c *gin.Context) {
	message := c.Query("message")
	if message == "" {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, message)
	}
}

func repeatMyParam(c *gin.Context) {
	message := c.Param("message")
	if message == "" {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, message)
	}
}

func repeatMyBody(c *gin.Context) {
	reqBody := c.Request.Body
	body, err := io.ReadAll(reqBody)
	if err != nil {
		log.Println(err)
	}
	unquotedBody, _ := strconv.Unquote(string(body))

	if unquotedBody == ""{
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, unquotedBody)
	}
}

func repeatMyHeader(c *gin.Context) {
	message := c.GetHeader("X-Message")
	if message == "" {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, message)
	}
}

func repeatMyCookie(c *gin.Context) {
	message, err := c.Cookie("message")
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, message)
	}
}