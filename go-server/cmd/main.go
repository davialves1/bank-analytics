package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://frontend"}
	server.Use(cors.New(corsConfig))

	server.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "time": time.Now()})
	})
	err := server.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println("Server failed to initiate", err)
	}

}
