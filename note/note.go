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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
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

// GIT_AUTHOR_DATE="Sat Nov 09 13:53:12 2024 +0200" GIT_COMMITTER_DATE="Sat Nov 09 13:53:12 2024 +0200" git commit -m "feat: adds struct metadata"
