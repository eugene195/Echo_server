package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	port        		int    `json:"port"`
	rootDir        		string `json:"rootDir"`
	indexFile 			string `json:"indexFile"`
	isInitialised 		bool
}

var config Config

func initConfig() {
	config = readConfig()
}

func getConfig() Config {
	return config
}

func readConfig() Config {
	str, err := ioutil.ReadFile("config.json")
	res := &Config{}
	if err != nil {
		res.isInitialised = false
		fmt.Println("error in file")
	} else {
		res.isInitialised = true
	}
	json.Unmarshal([]byte(str), &res)

	return *res
}