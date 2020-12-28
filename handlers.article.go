package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// render template with provided data
	render(c, "index.html", gin.H{
		"title":   "Home Page",
		"payload": articles,
	})
}

func getArticle(c *gin.Context) {
	// Check if the Article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the Article exists
		if article, err := getArticleByID(articleID); err == nil {
			render(c, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}
