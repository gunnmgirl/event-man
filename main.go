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

type displayer interface {
	Log()
}

type outputable interface {
	saver
	displayer
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello world")
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

	outputData(note)
	outputData(todo)
	printSomething(todo)
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

func outputData(data outputable) error {
	data.Log()
	err := saveData(data)
	if err != nil {
		return err
	}
	return nil
}

func printSomething(value any) {
	typedValue, ok := value.(int)
	fmt.Println("oke-", ok, typedValue)

	// switch value.(type) {
	// case int:
	// 	fmt.Println("integer value")
	// case float64:
	// 	fmt.Println("float value")
	// case string:
	// 	fmt.Println("string value")
	// default:
	// 	fmt.Println("default case")

	// }

	fmt.Println(value)
}
