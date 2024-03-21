package usecases

import (
	"log"
	"simple-hexagonal-arch-go-api/helpers"
	"simple-hexagonal-arch-go-api/internal/core/domain"
	"simple-hexagonal-arch-go-api/internal/core/ports"
)

type TodoUseCase struct {
	todoRepo ports.TodoRepository
}

func NewTodoUseCase(todoRepo ports.TodoRepository) ports.TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

// Create implements ports.TodoUseCase.
func (t *TodoUseCase) Create(title string, description string, done bool) (*domain.Todo, error) {
	todo := domain.NewTodo(helpers.RandomUUIDAsString(), title, description, done)

	_, err := t.todoRepo.Create(todo)
	if err != nil {
		log.Fatal("Error creating from repo", "todo")
		return nil, err
	}

	return todo, nil
}

// Done implements ports.TodoUseCase.
func (t *TodoUseCase) Done(id string) (*domain.Todo, error) {
	todo, err := t.Done(id)
	if err != nil {
		log.Fatal("Error done from repo")
		return nil, err
	}
	return todo, nil
}

// Get implements ports.TodoUseCase.
func (t *TodoUseCase) Get(id string) (*domain.Todo, error) {
	todo, err := t.todoRepo.Get(id)
	if err != nil {
		log.Fatal("Error getting from repo")
		return nil, err
	}
	return todo, nil
}

// List implements ports.TodoUseCase.
func (t *TodoUseCase) List() ([]domain.Todo, error) {
	todos, err := t.todoRepo.List()
	if err != nil {
		log.Fatal("Error listing from repo")
		return nil, err
	}

	return todos, nil
}
