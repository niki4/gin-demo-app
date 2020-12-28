package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
}

func initializeRoutes() {
	router.GET("/", getIndexHandler)
}
