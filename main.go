package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"strconv"

	"github.com/lskld/quippet/internal/model"
	"github.com/lskld/quippet/internal/storage"
	"github.com/lskld/quippet/internal/ui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: qpt [add|list]")
		return
	}

	switch os.Args[1] {
	case "add":
		err := addSnippet()
		if err != nil {
			fmt.Println("error: ", err)
		}
	case "list":
		err := listSnippets()
		if err != nil {
			fmt.Println("error: ", err)
		}
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
	
}

func addSnippet() error {
	//Saves the following variables: title, []tags, content

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose a title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	title = strings.TrimSpace(title)

	fmt.Print("Choose tags, seperate by comma (eg. banana, apple, pear): ")
	tagInput, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	tags := strings.Split(tagInput, ",")

	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}

	fmt.Print("Paste your snippet: ")
	content, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	content = strings.TrimSpace(content)

	newSnippet := model.Snippet {
		ID: strconv.FormatInt(time.Now().UnixNano(), 10),
		Title: title,
		Tags: tags,
		Content: content,
		CreatedAt: time.Now(),
	}

	snippets, err := storage.Load()
	if err != nil {
		return err
	}

	snippets = append(snippets, newSnippet)

	err = storage.Save(snippets)
	if err != nil {
		return err
	}

	return nil
}


func listSnippets() error {
	snippets, err := storage.Load()
	if err != nil {
		return err
	}

	err = ui.Run(snippets)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return nil
}