package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var router *gin.Engine

func getIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
}

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*") // cache templates in memory

	// register handlers
	router.GET("/", getIndexHandler)

	// start serving the app on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatal(router.Run())
}
