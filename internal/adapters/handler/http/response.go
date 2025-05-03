package http

import "github.com/gofiber/fiber/v2"

func ErrorResponse(ctx *fiber.Ctx, status int, msg string) error {
	return ctx.Status(status).JSON(fiber.Map{"success": false, "message": msg})
}

func SuccessResponse(ctx *fiber.Ctx, msg string, data any) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": msg, "data": data})
}
