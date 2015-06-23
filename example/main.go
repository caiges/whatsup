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
	Projects []map[string]string
	Slack    map[string]string
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

	contents := whatsup.GetContents(config.Projects)

	fmt.Printf("%s", contents)
}
