package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the HTML method of the Context to render a template
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
	// c.HTML(
	// 	http.StatusOK,
	// 	// Use the index.html template
	// 	"index.html",
	// 	// Pass the data that the page uses (in this case, 'title')
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
			// c.HTML(
			// 	http.StatusOK,
			// 	"article.html",
			// 	gin.H{
			// 		"title":   article.Title,
			// 		"payload": article,
			// 	},
			// )
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an Invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
