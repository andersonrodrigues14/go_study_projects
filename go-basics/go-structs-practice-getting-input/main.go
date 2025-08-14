package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	todo "example.com/note/todo2"
)

type Saver interface {
	Save() error
}

// type Displayer interface {
// 	Display()
// }

type outputtable interface {
	Saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	// printShomething(1.5)
	// printShomething(2)
	// printShomething("AAAAAA")
	// printShomething("TESTE")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	// todo.Display()
	// err = saveData(todo)
	err = outputData(todo)

	if err != nil {
		return
	}

	// userNote.Display()
	// err = saveData(userNote)
	err = outputData(userNote)

	if err != nil {
		return
	}
}

func printShomething(value interface{}) { // it's not necessary
	//typeVal, ok := value.(int) // ok - returns if value is int or not int / typeVal - returns the entire part os value

	//Search type with switch
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// default:
	// 	fmt.Println("Invalid Value: ", value)
	// }

	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}

	fmt.Println("saving the note suceeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
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
