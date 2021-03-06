package main

import (
	"github.com/gin-gonic/gin"
	_ "io/ioutil"
	"net/http"

	"fmt"
	"os"
)

const fileDir = "Upload"

// Define file type
type file struct {
	Name string `json:"name"`
	TimeStamp int `json:"ts"`
	Content []byte `json:"content"`
}

var recentFile file

// Handle the file upload
func getUploadFile(c *gin.Context) {
	var json file
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	recentFile = json

	go func() {
		path := fileDir + "/" + json.Name 
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0766)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		defer file.Close()

		file.Write(json.Content)
	}()

	c.JSON(http.StatusOK, gin.H{"msg": "Upload successfully!"})
}
