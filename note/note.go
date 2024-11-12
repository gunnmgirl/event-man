package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(Title, content string) (Note, error) {
	if Title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     Title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Log() {
	fmt.Printf("Note Title: %s\n", n.Title)
	fmt.Printf("Note content: %s\n", n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0o644)
}
