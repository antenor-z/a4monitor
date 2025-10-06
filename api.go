package main

import (
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type resultsStruct struct {
	AllSystemsOk bool     `json:"allSystemsOk"`
	Results      []Result `json:"results"`
}
type ToBeMonitored struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
}
type Config struct {
	Title         string          `json:"title"`
	ToBeMonitored []ToBeMonitored `json:"monitor"`
}

type Result struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

var results resultsStruct

func checkProcess(cfg *ToBeMonitored, results *resultsStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	status := runCheck(cfg.Cmd)
	if status == 0 {
		results.Results = append(results.Results, Result{Name: cfg.Name, Status: "operational"})
	} else {
		results.Results = append(results.Results, Result{Name: cfg.Name, Status: "failed"})
		results.AllSystemsOk = false
	}
}

func checkProcesses(wg *sync.WaitGroup) {
	config, err := loadConfig("config.json")
	if err != nil {
		panic("no config.json")
	}
	results.AllSystemsOk = true
	results.Results = nil
	for _, cfg := range config.ToBeMonitored {
		wg.Add(1)
		go checkProcess(&cfg, &results, wg)
	}
	wg.Wait()
	slices.SortFunc(results.Results, func(a, b Result) int {
		return strings.Compare(a.Name, b.Name)
	})
}

func CheckProcessTimer() {
	var wg sync.WaitGroup
	timer := time.NewTicker(10 * time.Second)
	defer timer.Stop()

	for range timer.C {
		checkProcesses(&wg)
	}
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, results)
}
