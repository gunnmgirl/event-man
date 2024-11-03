package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	note.Log()
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}
