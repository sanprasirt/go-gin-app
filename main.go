package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	route = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes servign HTML pages vary fast.
	route.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone function that will be used as route handlers

	// Initialize the routes
	initializeRoutes()

	// route.GET("/", func(c *gin.Context) {
	// 	// Call the HTML method of context to render a template
	// 	c.HTML(
	// 		// Set the HTML status to 200 OK
	// 		http.StatusOK,
	// 		// Use the index.html template
	// 		"index.html",
	// 		// Pass the data that the page uses (in this case, 'title')
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	// Start serving the application
	route.Run()
}

// Render one of the HTML, JSON or CSV based on the accept header of the request
// IF the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Response with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Response with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Response with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
