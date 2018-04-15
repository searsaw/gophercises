package adventure

import (
	"net/http"
)

func NewHttpAdventure(jsonData []byte) (*http.ServeMux, error) {
	story, err := parseJSONData(jsonData)
	if err != nil {
		return nil, err
	}
}
