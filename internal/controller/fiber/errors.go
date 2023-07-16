package fiber

import (
	"halyk-todo-list-api/internal/data/database"

	"github.com/gofiber/fiber/v2"
)

type errorMessage struct {
	Message string `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch err {
	case database.ErrNotFound:
		return c.Status(fiber.StatusNotFound).JSON(errorMessage{err.Error()})
	case database.ErrAlreadyExists:
		return c.Status(fiber.StatusConflict).JSON(errorMessage{err.Error()})
	case database.ErrDuplicatedKey:
		return c.Status(fiber.StatusConflict).JSON(errorMessage{err.Error()})
	case database.ErrUnsupported:
		return c.Status(fiber.StatusBadRequest).JSON(errorMessage{err.Error()})
	case database.ErrInvalidData:
		return c.Status(fiber.StatusBadRequest).JSON(errorMessage{err.Error()})

	default:
		return c.Status(fiber.StatusInternalServerError).JSON(errorMessage{err.Error()})
	}
}
