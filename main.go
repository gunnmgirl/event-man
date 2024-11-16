package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Print(err)
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	note.Log()
	todo.Log()

	err = saveData(note)
	if err != nil {
		return
	}

	err = saveData(todo)
	if err != nil {
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving failed")
		return err
	}
	fmt.Println("saving succeeded")
	return nil
}
