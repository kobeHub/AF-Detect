package main

import (
	"math/rand"
	_ "net/http"
	"github.com/gin-gonic/gin"
)

var (
	tryData = []string{"Just be a test",
	"So there is a lot fun woth golang",
	"vue is a simple and friendly framwork"}
	
	randData = []int{190, 890, 90, rand.Intn(200), rand.Intn(500)}
)


// Handle the request and render the index page
func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Atrial Filrillation Detection",
		"payload": "Welcome to the test site",
	})	
}

func getTryData(c *gin.Context) {
	// Crossed fileds request, set the request header
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	render(c, gin.H{
		"try_data": tryData,
		"rand_data": randData,
	})	
}

