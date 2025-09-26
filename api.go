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

func checkProcess(cfg *Config, results *resultsStruct) {
	status := runCheck(cfg.Cmd)
	if status == 0 {
		results.Results = append(results.Results, Result{Name: cfg.Name, Status: "operational"})
	} else {
		results.Results = append(results.Results, Result{Name: cfg.Name, Status: "failed"})
		results.AllSystemsOk = false
	}
}
func checkProcesses() {
	configs, err := loadConfig("config.json")
	if err != nil {
		panic("no config.json")
	}
	results.AllSystemsOk = true
	results.Results = nil
	for _, cfg := range configs {
		go checkProcess(&cfg, &results)
	}
}

func CheckProcessTimer() {
	timer := time.NewTicker(10 * time.Second)
	defer timer.Stop()

	for range timer.C {
		checkProcesses()
	}
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, results)
}
