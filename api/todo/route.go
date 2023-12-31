package todo

import (
	"go-web/core/todo"

	"github.com/gofiber/fiber/v2"
)

func Handle(app fiber.Router) {

	app.Get("/todo-form", todo.HandleGetFormTemplate)

	app.Get("/todos/:id", todo.HandleGetItemTemplate)

	app.Post("/todo", todo.HandlePostTodoItem)

	app.Delete("/todos/:id", todo.HandleDeleteTodoItem)
}
