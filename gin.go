package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/events", events)
	router.Run(":5000")
}

func events(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	resp := map[string]string{"hello": "world"}
	c.JSON(http.StatusOK, resp)
	/*post_gwid := c.PostForm("name")
	fmt.Println(post_gwid)*/
}
