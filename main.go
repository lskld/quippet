package main

import (
	"fmt"

	"github.com/lskld/quippet/internal/storage"
)

func main() {
	fmt.Println("qpt starting")
}

func listSnippets() error {
	snippets, err := storage.Load()
	if err != nil {
		return err
	}

	for _, value := range snippets {
		fmt.Printf("%s | %s | %v\n", value.Id, value.Title, value.Tags)
	}

	return nil
}