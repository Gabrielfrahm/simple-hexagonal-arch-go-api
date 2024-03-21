package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"simple-hexagonal-arch-go-api/helpers"
	"simple-hexagonal-arch-go-api/internal/core/ports"
	usecases "simple-hexagonal-arch-go-api/internal/core/usecases"
	handlerTodo "simple-hexagonal-arch-go-api/internal/handles/todo"
	repoTodo "simple-hexagonal-arch-go-api/internal/repositories/todo"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/joho/godotenv"
)

var (
	repo    string
	binding string
)

func init() {
	flag.StringVar(&repo, "repo", "postgres", "Mongo or Postgres")
	flag.StringVar(&binding, "httpbind", ":3333", "address/port to bind listen socket")

	flag.Parse()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var todoRepo ports.TodoRepository
	if repo == "postgres" {
		todoRepo = startPostgresRepo()
	}

	todoUseCase := usecases.NewTodoUseCase(todoRepo)

	ws := new(restful.WebService)
	ws = ws.Path("/api")
	handlerTodo.NewHandler(todoUseCase, ws)
	restful.Add(ws)

	fmt.Println("Listening...")

	log.Panic(http.ListenAndServe(binding, nil))
}

func startPostgresRepo() ports.TodoRepository {
	return repoTodo.NewTodoPostgresRepo(helpers.StartPostgresDb())
}
