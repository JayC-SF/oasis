package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Redirect the home path to the app path
func redirectBasePath(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/app")
}

// Function takes care of configuring the the base
func serverConfiguration(router *gin.Engine) {
	router.Static("/app", filepath.Join("..", "frontend", "dist"))
	router.GET("/", redirectBasePath)
}

func main() {
	router := gin.Default()
	// Configure the server redirects and static files
	serverConfiguration(router)

	router.Run(":8080")
}
