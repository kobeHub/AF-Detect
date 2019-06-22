package main

import (
	"strconv"
	"fmt"
	"os/exec"
	"bytes"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/buger/jsonparser"
)

// The Enumerate of labels
type record struct {
	FileName string
	TimeStamp int
	AF float64 
	NO float64
	UN float64
}


var cache map[int]record

// Handle the detect 
func getDetectResult(c *gin.Context) {
	ops := c.PostForm("ops")

	fmt.Println(ops)
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
	return

	ts := strconv.Itoa(recentFile.TimeStamp)
	command := "./predict Upload/" + ts + "-" + recentFile.Name
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

	detect_res := record{
		FileName: recentFile.Name,
		TimeStamp: recentFile.TimeStamp,
		AF: af,
		NO: no,
		UN: un,
	}	
	
	cache[recentFile.TimeStamp] = detect_res
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
