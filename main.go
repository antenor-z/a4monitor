package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	go CheckProcessTimer()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/api/status", getStatus)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":5300")
}
