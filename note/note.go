package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (n Note) Log() {
	fmt.Printf("Note title: %s\n", n.title)
	fmt.Printf("Note content: %s\n", n.content)
}
