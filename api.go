package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type resultsStruct struct {
	AllSystemsOk bool     `json:"allSystemsOk"`
	Results      []Result `json:"results"`
}

var results resultsStruct

func checkProcess() {
	configs, err := loadConfig("config.json")
	if err != nil {
		panic("no config.json")
	}
	results.AllSystemsOk = true
	clear(results.Results)
	for _, cfg := range configs {
		status := runCheck(cfg.Cmd)
		if status == 0 {
			results.Results = append(results.Results, Result{Name: cfg.Name, Status: "operational"})
		} else {
			results.Results = append(results.Results, Result{Name: cfg.Name, Status: "failed"})
			results.AllSystemsOk = false
		}
	}
	print("**")
}

func CheckProcessTimer() {
	timer := time.NewTicker(time.Second * 10)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			go func() {
				checkProcess()
			}()
		}
	}
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, results)
}
