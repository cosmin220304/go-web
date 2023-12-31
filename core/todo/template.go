package todo

import (
	model "go-web/model"
	util "go-web/util"

	"github.com/gofiber/fiber/v2"
)

func TodoFormTemplate(todos []model.Todo, ctx *fiber.Ctx) (string, bool) {
	return util.GetTemplate(ctx, todoForm(todos))
}

func TodoListTemplate(todos []model.Todo, ctx *fiber.Ctx) (string, bool) {
	return util.GetTemplate(ctx, todosList(todos))
}

func TodoItemTemplate(todo model.Todo, ctx *fiber.Ctx) (string, bool) {
	return util.GetTemplate(ctx, todoItem(todo))
}
