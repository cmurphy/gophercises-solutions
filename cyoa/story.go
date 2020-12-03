package cyoa

import (
	"encoding/json"
	"io/ioutil"
)

// Story is a map of keys to chapters
type Story map[string]Chapter

// Chapter is a segment of a story
type Chapter struct {
	Title      string
	Paragraphs []string `json:"story"`
	Options    []Option
}

// Option is a choice for the reader
type Option struct {
	Text    string
	Chapter string `json:"arc"`
}

// ReadStory reads a JSON file containing a story and renders it into a map
func ReadStory(file string) (Story, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	story := make(Story)
	err = json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
