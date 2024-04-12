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

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("any")

	title, content := getNoteData()
	todoText := getUser("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}
}

func printSomething(value any) {
	intVal, ok := value.(int)

	if !ok {
		fmt.Println("Interger: ", intVal)
		return
	}

	floatVal, ok := value.(int)

	if !ok {
		fmt.Println("Float: ", floatVal)
		return
	}

	stringVal, ok := value.(int)

	if !ok {
		fmt.Println("String: ", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Interger: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println(value)
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	data.Save()

	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUser("Note title: ")
	content := getUser("Note Content: ")

	return title, content
}

func getUser(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
