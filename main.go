package main

import (
	"context"
	"go-web/api/todo"
	"go-web/data"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var fiberLambda *fiberadapter.FiberLambda

func setupApp() *fiber.App {
	data.ConnectDb()

	app := fiber.New()

	app.Static("/", "./public")
	todo.Handle(app)

	return app
}

func main() {
	if os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		fiberLambda = fiberadapter.New(setupApp())
		lambda.Start(HandleRequest)
	} else {
		app := setupApp()

		// Handle the WebSocket connection, used for live updates in development mode
		if os.Getenv("APP_ENV") == "" {
			app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
				ctx.ReadMessage()
			}))
		}

		app.Listen(":3000")
	}
}

// Handler will deal with Fiber working with Lambda
func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	res, error := fiberLambda.ProxyWithContext(ctx, req)
	if res.Headers == nil {
		res.Headers = make(map[string]string)
	}
	res.Headers["Content-Type"] = "text/html"
	return res, error
}
