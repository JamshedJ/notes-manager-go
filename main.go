package main

import "fmt"

var notes []Note
var nextID int = 1

type Note struct {
	ID      int
	Title   string
	Content string
}

func createNote() {
	var title, content string

	fmt.Println("Enter the title of the note: ")
	fmt.Scanln(&title)

	fmt.Println("Enter the content of the note: ")
	fmt.Scanln(&content)

	notes = append(notes, Note{
		ID:      nextID,
		Title:   title,
		Content: content,
	})

	nextID++
}

func getNotes() {
	fmt.Println("All notes")

	if len(notes) == 0 {
		fmt.Println("You have no notes")
		return
	}

	for _, note := range notes {
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", note.ID, note.Title, note.Content)
	}
}

func getNoteByID() {
	var id int
	fmt.Println("Enter the ID of the note: ")
	fmt.Scanln(&id)

	for _, note := range notes {
		if note.ID == id {
			fmt.Printf("ID: %d, Title: %s, Content: %s\n", note.ID, note.Title, note.Content)
			return
		} else {
			fmt.Println("Note not found")
		}
	}
}

func updateNote() {
	var id int
	fmt.Println("Enter the ID of the note: ")
	fmt.Scanln(&id)

	for i, note := range notes {
		if note.ID == id {
			var title, content string

			fmt.Println("Enter the title of the note: ")
			fmt.Scanln(&title)

			fmt.Println("Enter the content of the note: ")
			fmt.Scanln(&content)

			notes[i] = Note{
				ID:      id,
				Title:   title,
				Content: content,
			}

			fmt.Println("Note updated")
			return
		} else {
			fmt.Println("Note not found")
		}
	}
}

func deleteNote() {
	var id int 
	fmt.Println("Enter the ID of the note: ")
	fmt.Scanln(&id)

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			fmt.Println("Note deleted")
			return
		} else {
			fmt.Println("Note not found")
		}
	}
}

func main() {
	for {
		fmt.Println("\nConsole Notes manager")
		fmt.Println("1. Create a note")
		fmt.Println("2. List all notes")
		fmt.Println("3. Get a note by ID")
		fmt.Println("4. Update note")
		fmt.Println("5. Delete note")
		fmt.Println("6. Exit")
		fmt.Println("\nEnter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createNote()
		case 2:
			getNotes()
		case 3:
			getNoteByID()
		case 4:
			updateNote()
		case 5:
			deleteNote()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
