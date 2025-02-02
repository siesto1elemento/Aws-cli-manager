package state

import (
	"encoding/json"
	"os"
)

type Resource struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Status string                 `json:"status"`
	Config map[string]interface{} `json:"config"`
}

func LoadState(filename string) (map[string]Resource, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var state map[string]Resource
	err = json.Unmarshal(file, &state)
	return state, err
}

func SaveState(filename string, state map[string]Resource) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
