package util

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func GetTemplate(ctx *fiber.Ctx, templComponent templ.Component) (string, bool) {
	htmlString, err := templ.ToGoHTML(ctx.Context(), templComponent)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return string(htmlString), true
}
