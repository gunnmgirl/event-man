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

	err = todo.Save()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("saving todo success")

	err = note.Save()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("saving note success")
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
