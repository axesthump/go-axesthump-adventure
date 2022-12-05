package arc

import (
	"encoding/json"
	"os"
)

type Option struct {
	Text    string `json:"text"`
	ArcName string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func GetArcs() (map[string]Arc, error) {
	fileData, err := os.ReadFile("gopher.json")
	if err != nil {
		return nil, err
	}
	arcs := make(map[string]Arc)

	err = json.Unmarshal(fileData, &arcs)
	if err != nil {
		return nil, err
	}

	return arcs, nil
}
