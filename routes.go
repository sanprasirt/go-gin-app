package main

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user ro not
	// route.Use(setUserStatus())

	// Handle the index route
	route.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := route.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show login page
		// Ensure that the user is not logged in by using the middleware
		// userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by useing the middleware
		// userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		// userRoutes.GET("/logout", ensureLoggedIn(), logout)

		userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", register)
	}

	// Handle GET requests at /article/view/some_article_id
	route.GET("/article/view/:article_id", getArticle)

	// Group article related routes together
	//articleRoutes := route.Group("/article")
	//{
	// articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
	// articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	//}
}
