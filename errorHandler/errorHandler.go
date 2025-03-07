package errorHandler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var response Error
	ctx.Set("Content-Type", "application/json")

	if e, ok := err.(*NotFoundError); ok {
		code := http.StatusNotFound
		ctx.Status(code)
		response = Error{
			Message:    e.Error(),
			Code:       code,
			StatusText: http.StatusText(code),
		}
		ctx.JSON(response)

		return err
	}

	code := http.StatusBadRequest
	ctx.Status(code)
	response = Error{
		Message:    err.Error(),
		Code:       code,
		StatusText: http.StatusText(code),
	}

	ctx.JSON(response)

	return err
}
