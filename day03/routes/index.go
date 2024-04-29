package routes

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) error {
    applyHealth(r)
    applyHello(r)
    applyRepeat(r)
    applyPalindrome(r)
    return nil
}