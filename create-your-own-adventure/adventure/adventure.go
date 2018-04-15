package adventure

import (
	"encoding/json"
)

type ArcOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string       `json:"title"`
	Story   []string     `json:"story"`
	Options []ArcOptions `json:"options"`
}

func parseJSONData(jsonData []byte) (map[string]StoryArc, error) {
	var story map[string]StoryArc
	if err := json.Unmarshal(jsonData, &story); err != nil {
		return nil, err
	}
	return story, nil
}
