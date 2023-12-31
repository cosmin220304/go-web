package todo

import (
	"go-web/model"

	"github.com/gofiber/fiber/v2"
)

func HandleGetFormTemplate(ctx *fiber.Ctx) error {
	todos := GetTodos()
	template, ok := TodoFormTemplate(todos, ctx)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return ctx.SendString(template)
}

func HandleGetItemTemplate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	item, ok := GetTodoById(id)
	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo item not found"})
	}

	template, ok := TodoItemTemplate(item, ctx)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return ctx.SendString(template)
}

func HandlePostTodoItem(ctx *fiber.Ctx) error {
	newTodo := model.Todo{Name: ctx.FormValue("name")}
	if newTodo.Name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	createdTodo, ok := CreateTodoItem(newTodo)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create todo item"})
	}

	template, ok := TodoItemTemplate(createdTodo, ctx)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return ctx.SendString(template)
}

func HandleDeleteTodoItem(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if ok := DeleteTodoById(id); !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo item not found"})
	}
	return ctx.Status(fiber.StatusOK).Send(nil)
}
