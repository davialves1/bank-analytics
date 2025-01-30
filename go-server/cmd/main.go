package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	err := server.Run("localhost:8080")
	if err != nil {
		fmt.Println("Server failed to initiate", err)
	}

}
