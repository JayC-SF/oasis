package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.NoRoute(static.Serve("/", static.LocalFile("../frontend/dist", true)))
	router.Run(":9090")
}
