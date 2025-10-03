package main

var config *Config

func getTitle() string {
	config, err := loadConfig("config.json")
	if err != nil {
		panic("no config.json")
	}
	return config.Title
}
