package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
}

type Result struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/api/status", check)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":8080")
}
