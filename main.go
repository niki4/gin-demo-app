package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*") // cache templates in memory

	// register handlers
	initializeRoutes()

	// start serving the app on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatal(router.Run())
}
