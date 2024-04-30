package routes

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func applyRepeat(r *gin.Engine) {
	r.GET("/repeat-my-query", repeatMyQuery)
	r.GET("/repeat-my-param/:message", repeatMyParam)
	r.POST("/repeat-my-body", repeatMyBody)
	r.GET("/repeat-my-header", repeatMyHeader)
	r.GET("/repeat-my-cookie", repeatMyCookie)
	r.GET("/repeat-all-my-queries", repeatAllMyQueries)
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

func repeatAllMyQueries(c *gin.Context) {
	type Query struct {
		Key string `json:"key"`
		Value string `json:"value"`
	}

	queries := c.Request.URL.Query()
	if len(queries) == 0 {
		c.Status(http.StatusBadRequest)
	} else {
		var queriesList []Query
		for key, value := range queries {
			queriesList = append(queriesList, Query{Key: key, Value: value[0]})
		}
		c.JSON(http.StatusOK, queriesList)
	}
}