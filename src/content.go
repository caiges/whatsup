package main

import (
	"net/http"
	"strconv"
)

type Content struct {
	Project string
	Version string
}

func GetContents(urls []map[string]string) []Content {
	var contents []Content

	for _, url := range urls {
		content := Content{}
		content.Project = url["project"]
		content.Version = GetVersion(url["url"])
		contents = append(contents, content)
	}

	return contents
}

func GetVersion(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		return "Unavailable"
	}

	if resp.StatusCode == 200 {
		version, err := ParseVersion(resp.Body)

		if err != nil {
			return "Invalid Version Format"
		}

		return version.Version
	} else {
		return strconv.Itoa(resp.StatusCode)
	}
}
