package todo

import (
	"simple-hexagonal-arch-go-api/internal/core/ports"

	restful "github.com/emicklei/go-restful/v3"
)

type TodoHandle struct {
	todoUseCase ports.TodoUseCase
}

func NewHandler(todoUseCase ports.TodoUseCase, ws *restful.WebService) *TodoHandle {
	handler := &TodoHandle{
		todoUseCase: todoUseCase,
	}

	ws.Route(ws.GET("/todo/{id}").To(handler.Get).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/todo").To(handler.List).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.POST("/todo").To(handler.Create).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	// ws.Route(ws.POST("/todo/{id}").To(handler.Done).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	return handler
}

func (tdh *TodoHandle) Get(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")

	result, err := tdh.todoUseCase.Get(id)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todo *Todo = &Todo{}

	todo.FromDomain(result)
	resp.WriteAsJson(todo)
}

func (tdh *TodoHandle) Create(req *restful.Request, resp *restful.Response) {
	var data = new(Todo)
	if err := req.ReadEntity(data); err != nil {
		resp.WriteError(500, err)
		return
	}
	result, err := tdh.todoUseCase.Create(data.Title, data.Description, data.Done)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todo Todo = Todo{}
	todo.FromDomain(result)
	resp.WriteAsJson(todo)
}

func (tdh *TodoHandle) List(req *restful.Request, resp *restful.Response) {
	result, err := tdh.todoUseCase.List()
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todos TodoList = TodoList{}

	todos = todos.FromDomain(result)
	resp.WriteAsJson(todos)
}

// func (tdh *TodoHandle) Done(req *restful.Request, resp *restful.Response) {
// 	id := req.PathParameter("id")

// 	result, err := tdh.todoUseCase.Get(id)
// 	if err != nil {
// 		resp.WriteError(500, err)
// 		return
// 	}

// 	var todo *Todo = &Todo{}

// 	todo.FromDomain(result)
// 	resp.WriteAsJson(todo)
// }
