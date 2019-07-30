package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// The global render variable
var router *gin.Engine
const PORT = ":8080"

/*old version*/
//func render(c *gin.Context, data gin.H, templateName string) {
//	switch c.Request.Header.Get("Accept") {
//	case "application/json":
//		c.JSON(http.StatusOK, data)
//	case "application/xml":
//		c.XML(http.StatusOK, data)
//	default:
//		c.HTML(http.StatusOK, templateName, data)
//	}
//}


// Backend only give json ans xml format
func render(c *gin.Context, data gin.H) {
	switch c.Request.Header.Get("Accept") {
		case "application/xml":
			c.XML(http.StatusOK, data)
		default:
			c.JSON(http.StatusOK, data)
	}
}

func main() {
	router = gin.Default()
	
	// Get redis instance
	// redisPool = RedisPolling()

	// Load templates for all pages
	// router.LoadHTMLGlob("templates/*")

	// Define the router 
	initRouters()
	
	fmt.Println("Serve at port:", PORT)	
	router.Run(PORT)
}
