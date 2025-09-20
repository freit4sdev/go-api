package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})
	
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}