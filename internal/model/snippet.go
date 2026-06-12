package model

import "time"

type Snippet struct {
	ID      string
	Title   string
	Tags    []string
	Content string
	CreatedAt time.Time
}