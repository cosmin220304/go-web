package todos

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

var todos = []string{"Buy milk", "Buy eggs", "Buy bread"}

func Handle(router fiber.Router) {
	router.Get("/todos", func(ctx *fiber.Ctx) error {
		return renderTodoList(ctx)
	})

	router.Post("/todo", func(ctx *fiber.Ctx) error {
		newTodo := ctx.FormValue("todo")
		todos = append(todos, strings.Clone(newTodo))
		return renderTodoList(ctx)
	})

	router.Delete("/todos/:id", func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}
		todos = append(todos[:id], todos[id+1:]...)
		return renderTodoList(ctx)
	})
}

func renderTodoList(ctx *fiber.Ctx) error {
	htmlString, err := templ.ToGoHTML(ctx.Context(), todosList(todos))
	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendString(string(htmlString))
}
