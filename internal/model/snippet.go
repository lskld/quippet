package model

import "time"

type Snippet struct {
	Id      string
	Title   string
	Tags    []string
	Content string
	CreatedAt time.Time
}