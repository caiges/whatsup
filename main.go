package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Urls []map[string]string
}

type Version struct {
	Version string `json:"version"`
}

func ParseVersion(data io.Reader) (Version, error) {
	decoder := json.NewDecoder(data)
	v := Version{}
	err := decoder.Decode(&v)

	return v, err
}

func main() {
	configFile, err := os.Open("config.json")
	checkErr(err)

	decoder := json.NewDecoder(configFile)
	config := Config{}
	err = decoder.Decode(&config)
	checkErr(err)

	for _, url := range config.Urls {
		resp, err := http.Get(url["url"])
		checkErr(err)

		if resp.StatusCode == 200 {
			version, err := ParseVersion(resp.Body)
			checkErr(err)

			fmt.Printf("%s -- %s\n", url["name"], version.Version)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
