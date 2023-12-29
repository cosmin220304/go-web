package main

import (
	todos "go-web/todos"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./todos/index.html")
	})

	// Handle the WebSocket connection, used for live updates in development mode
	if os.Getenv("APP_ENV") == "development" {
		app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
			ctx.ReadMessage()
		}))
	}

	todos.Handle(app)

	panic(app.Listen(":3000"))
}
