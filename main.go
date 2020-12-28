package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode. Comment the line for local debug.
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()
	router.LoadHTMLGlob("templates/*") // cache templates in memory

	// register handlers
	initializeRoutes()

	// start serving the app on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatal(router.Run())
}

// Response in either HTML, JSON or CSV depending on request "Accept" header
// Fall back on HTML if the header not specified - templateName required for this.
func render(c *gin.Context, templateName string, data gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
