package whatsup

import (
	"encoding/json"
	"io"
)

// Version is an internal structure for
type Version struct {
	Version string `json:"version"`
}

func ParseVersion(data io.Reader) (Version, error) {
	decoder := json.NewDecoder(data)
	v := Version{}
	err := decoder.Decode(&v)

	return v, err
}
