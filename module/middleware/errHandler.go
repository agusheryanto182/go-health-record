package middleware

import (
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := ctx.Response().StatusCode()
	if customError, ok := err.(response.CustomError); ok {
		code = customError.Status()
		return ctx.Status(code).JSON(fiber.Map{
			"message": customError.Error(),
		})
	} else if code < 400 {
		code = fiber.StatusInternalServerError
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
