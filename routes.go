package main


// Define URL and its func handler
func initRouters() {
	// Index page
	router.GET("/", showIndexPage)
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Vue try data
	router.GET("/api/try", getTryData)
	router.POST("/api/upload", getUploadFile)
	router.POST("/api/detect", getDetectResult)
}
