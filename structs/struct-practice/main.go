package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
	}	

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err:= reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	// Remove the newline character from the end of the input
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // For Windows compatibility
	
	return text
}
