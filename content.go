package whatsup

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

var timeout = time.Duration(5 * time.Second)

type Content struct {
	Project string
	Env     string
	Mode    string
	Version string
}

func GetContents(urls []map[string]string) []Content {
	var contents []Content

	for _, url := range urls {
		content := Content{}
		content.Project = url["project"]
		content.Env = url["env"]
		content.Mode = url["mode"]
		content.Version = GetVersion(url["url"])
		contents = append(contents, content)
	}

	return contents
}

func GetVersion(url string) string {
	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}

	resp, err := client.Get(url)

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

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}
