package routes

import (
	"SoftwareGoDay3/middlewares"
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
	router.GET("/health", health)
	router.GET("/repeat-all-my-queries", repeatAllMyQueries)
	// router.POST("/are-these-palindromes", areThesePalindromes)
	router.POST("/are-these-palindromes", middlewares.CheckPalindrome(), areThesePalindromes)
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

func health(c *gin.Context) {
	c.Status(http.StatusOK)
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

// Structure for our returned JSON
type PalindromeResponse struct {
    Input  string `json:"input"`
    Result bool   `json:"result"`
}  
// Helper function to check a single string
func isPalindrome(input string) bool {
    size := len(input)
    stop := size / 2
    for i := 0; i < stop; i++ {
        if input[i] != input[size-i-1] {
            return false
        }
    }
    return true
}  
// Main function
func areThesePalindromes(c *gin.Context) {
		inputs := c.MustGet("inputs").([]string)

    palindromes := make([]PalindromeResponse, len(inputs))
    for idx, input := range inputs {
        palindromes[idx] = PalindromeResponse{Input: input, Result: isPalindrome(input)}
    }
    c.JSON(http.StatusOK, palindromes)
}