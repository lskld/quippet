package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/lskld/quippet/internal/model"
)

func Load() ([]model.Snippet, error) {
	path, err := buildFilePath()
	if err != nil {
		return nil, err
	}

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
	path, err := buildFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(snippets)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func buildFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, ".quippet", "snippets.json")

	return path, nil
}