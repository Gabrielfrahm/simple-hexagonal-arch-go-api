package ports

import "simple-hexagonal-arch-go-api/internal/core/domain"

type TodoRepository interface {
	Get(id string) (*domain.Todo, error)
	List() ([]domain.Todo, error)
	Create(todo *domain.Todo) (*domain.Todo, error)
	Done(id string) (*domain.Todo, error)
}
