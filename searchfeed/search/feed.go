package search

import (
	"encoding/json"
	"os"
)

// lowercase constant is unexported to other packages
const dataFile = "data/data.json"

// Holds the metadata of feed
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// reads and unmarshals the data feed file
func RetrieveFeeds() ([]*Feed, error) {
	// open the file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// ensure file is to be closed
	defer file.Close()

	// decode the file into a slice of pointers to feed values
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err // caller to check for decoding errors
}
