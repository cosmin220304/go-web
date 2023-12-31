package main

import (
	"go-web/api/todo"
	"go-web/data"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	data.ConnectDb()

	app := fiber.New()
	app.Static("/", "./public")

	todo.Handle(app)

	// Handle the WebSocket connection, used for live updates in development mode
	if os.Getenv("APP_ENV") == "" {
		app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
			ctx.ReadMessage()
		}))
	}

	panic(app.Listen(":3000"))
}
