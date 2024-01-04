package todo

import (
	"go-web/model"
	util "go-web/util"

	"github.com/gofiber/fiber/v2"
)

func HandleGetFormTemplate(ctx *fiber.Ctx) error {
	todos := GetTodos()
	template, ok := util.GetTemplate(ctx, TodoForm(todos))
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return ctx.SendString(template)
}

func HandleGetItemTemplate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todoItem, ok := GetTodoById(id)
	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo item not found"})
	}

	template, ok := util.GetTemplate(ctx, TodoItem(todoItem))
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

	template, ok := util.GetTemplate(ctx, TodoItem(createdTodo))
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
