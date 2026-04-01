package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", note.title, note.content)
}

func (note Note) Save() {
	// Implementation for saving the note
	filename := strings.ReplaceAll(note.title, " ", "_")
	filename = strings.ToLower(filename)
	os.WriteFile(filename)
}


func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}
	
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}