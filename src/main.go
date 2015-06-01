package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Config is an internal representation of the config file
type Config struct {
	Urls  []map[string]string
	Slack map[string]string
}

func main() {
	configFile, err := os.Open("config.json")

	if err != nil {
		log.Fatal("Could not open config")
	}

	decoder := json.NewDecoder(configFile)
	config := Config{}
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal("Could not parse config")
	}

	contents := GetContents(config.Urls)

	fmt.Printf("%s", contents)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
