package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	Urls	[]string `json:"urls"` 
}

func main() {
	configFile, err := os.Open("config.json")
	checkErr(err)

	decoder := json.NewDecoder(configFile)
	config := Config{}
	err = decoder.Decode(&config)
	checkErr(err)

	fmt.Println(config.Urls)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}