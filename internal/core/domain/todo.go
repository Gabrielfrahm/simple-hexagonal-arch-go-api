package domain

import "fmt"

type Todo struct {
	ID          string
	Title       string
	Description string
	Done        bool
}

func NewTodo(id string, title, description string, done bool) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Done:        done,
	}
}

func (t *Todo) String() string {
	return fmt.Sprintf("%s - %s - %v", t.Title, t.Description, t.Done)
}

func (t *Todo) MakeDone() {
	t.Done = !t.Done
}
