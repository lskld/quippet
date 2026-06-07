package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/lskld/quippet/internal/model"
)

func Load() ([]model.Snippet, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}
	
	path := filepath.Join(home, ".quippet", "snippets.json")

	data, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			return []model.Snippet{}, nil
		}
		return nil, err
	}

	var snippets []model.Snippet
	err = json.Unmarshal(data, &snippets)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func Save(snippets []model.Snippet) error {
	return nil
}