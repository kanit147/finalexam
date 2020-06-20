package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizationMdw(c *gin.Context) {
	log.Println("start authorizationMdw")
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access denied"})
		c.Abort()
		return
	}

	c.Next()

	log.Println("end authorizationMdw")

}
