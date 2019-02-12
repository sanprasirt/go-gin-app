package main

func initializeRoutes() {
	// Handle the index route
	route.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	route.GET("/article/view/:article_id", getArticle)
}
