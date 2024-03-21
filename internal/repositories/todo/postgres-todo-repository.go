package todo

import (
	"database/sql"
	"fmt"
	"simple-hexagonal-arch-go-api/internal/core/domain"
	"simple-hexagonal-arch-go-api/internal/core/ports"
)

type todoPostgres struct {
	ID          string
	Title       string
	Description string
	Done        bool
}

type todoListPostgres []todoPostgres

func (p *todoPostgres) ToDomain() *domain.Todo {
	return &domain.Todo{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Done:        p.Done,
	}
}

func (p *todoPostgres) FromDomain(todo *domain.Todo) {
	if p == nil {
		p = &todoPostgres{}
	}

	p.ID = todo.ID
	p.Title = todo.Title
	p.Description = todo.Description
	p.Done = todo.Done
}

func (p todoListPostgres) ToDomain() []domain.Todo {
	todos := make([]domain.Todo, len(p))
	for k, td := range p {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos
}

type todoPostgresRepo struct {
	db *sql.DB
}

func NewTodoPostgresRepo(db *sql.DB) ports.TodoRepository {
	return &todoPostgresRepo{
		db: db,
	}
}

// Create implements ports.TodoRepository.
func (t *todoPostgresRepo) Create(todo *domain.Todo) (*domain.Todo, error) {
	sqlS := "INSERT INTO todo (id, title, description, done) VALUES ($1, $2, $3, $4)"

	_, err := t.db.Exec(sqlS, todo.ID, todo.Title, todo.Description, todo.Done)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Done implements ports.TodoRepository.
func (t *todoPostgresRepo) Done(id string) (*domain.Todo, error) {
	panic("unimplemented")
}

// Get implements ports.TodoRepository.
func (t *todoPostgresRepo) Get(id string) (*domain.Todo, error) {
	var todo todoPostgres = todoPostgres{}
	sqsS := fmt.Sprintf("SELECT id, title, description, done FROM todo WHERE id = '%s'", id)

	result := t.db.QueryRow(sqsS)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done); err != nil {
		return nil, err
	}

	return todo.ToDomain(), nil
}

// List implements ports.TodoRepository.
func (t *todoPostgresRepo) List() ([]domain.Todo, error) {
	var todos todoListPostgres
	sqsS := "SELECT id, title, description FROM todo"

	result, err := t.db.Query(sqsS)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		todo := todoPostgres{}

		if err := result.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos.ToDomain(), nil
}
