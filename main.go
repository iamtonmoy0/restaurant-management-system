package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	r := gin.New()
	r.Use(gin.Logger())

	r.Run(":" + port)
}
