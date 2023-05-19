package json

import "github.com/gofiber/fiber/v2"

func (respdata ReturnData) WriteToBody(ctx *fiber.Ctx) error {
	return ctx.Status(respdata.Code).JSON(respdata)
}
