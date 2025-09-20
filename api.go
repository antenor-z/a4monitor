package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(c *gin.Context) {
	configs, err := loadConfig("config.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var results []Result
	allSystemsOk := true
	for _, cfg := range configs {
		status := runCheck(cfg.Cmd)
		if status == 0 {
			results = append(results, Result{Name: cfg.Name, Status: "operational"})
		} else {
			results = append(results, Result{Name: cfg.Name, Status: "failed"})
			allSystemsOk = false
		}
	}

	c.JSON(http.StatusOK, gin.H{"allSystemsOk": allSystemsOk, "results": results})
}
