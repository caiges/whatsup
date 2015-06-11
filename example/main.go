package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caiges/whatsup"
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
	config := whatsup.Config{}
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal("Could not parse config")
	}

	contents := whatsup.GetContents(config.Urls)

	fmt.Printf("%s", contents)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
