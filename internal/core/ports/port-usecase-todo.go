package ports

import "simple-hexagonal-arch-go-api/internal/core/domain"

type TodoUseCase interface {
	Get(id string) (*domain.Todo, error)
	List() ([]domain.Todo, error)
	Create(title, description string, done bool) (*domain.Todo, error)
	Done(id string) (*domain.Todo, error)
}
