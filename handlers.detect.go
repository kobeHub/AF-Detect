package main

import (
	"fmt"
	"os/exec"
	"bytes"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/buger/jsonparser"
)

// The Enumerate of labels
type record struct {
	Name string    `json:"name"` 
}

// Handle the detect 
func getDetectResult(c *gin.Context) {
	var json record
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command := "./predict Upload/" + json.Name + ".mat"
	res, err := detect(command)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	data := []byte(res)
	af, _ := jsonparser.GetFloat(data, "target_pred", "A")
	no, _ := jsonparser.GetFloat(data, "target_pred", "NO")
	un, _ := jsonparser.GetFloat(data, "uncertainty")

	c.JSON(http.StatusOK, gin.H{
		"AF": af,
		"NO": no,
		"UN": un,
	})
}

// Call the `predict` to get detect results
func detect(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	result := out.String()
	return result, err
}
