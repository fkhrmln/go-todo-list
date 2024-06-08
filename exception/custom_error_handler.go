package exception

import (
	"go-todo-list/dto/response"

	"github.com/gofiber/fiber/v2"
)

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *ValidationError:
		response := response.Response{
			Status:  "Bad Request",
			Code:    fiber.StatusBadRequest,
			Message: e.Error(),
			Data:    nil,
		}

		return c.Status(400).JSON(response)
	case *BadRequestError:
		response := response.Response{
			Status:  "Bad Request",
			Code:    fiber.StatusBadRequest,
			Message: e.Message,
			Data:    nil,
		}

		return c.Status(400).JSON(response)
	case *UnauthorizedError:
		response := response.Response{
			Status:  "Unauthorized",
			Code:    fiber.StatusUnauthorized,
			Message: e.Message,
			Data:    nil,
		}

		return c.Status(401).JSON(response)
	case *NotFoundError:
		response := response.Response{
			Status:  "Not Found",
			Code:    fiber.StatusNotFound,
			Message: e.Message,
			Data:    nil,
		}

		return c.Status(404).JSON(response)
	default:
		response := response.Response{
			Status:  "Internal Server Error",
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}

		return c.Status(500).JSON(response)
	}
}

func CustomTokenErrorHandler(c *fiber.Ctx, err error) error {
	switch err.Error() {
	case "missing or malformed JWT":
		response := response.Response{
			Status:  "Bad Request",
			Code:    fiber.StatusBadRequest,
			Message: c.Path() + " Require Token",
			Data:    nil,
		}

		return c.Status(400).JSON(response)
	case "token is expired":
		response := response.Response{
			Status:  "Unauthorized",
			Code:    fiber.StatusUnauthorized,
			Message: "Token Expired",
			Data:    nil,
		}

		return c.Status(401).JSON(response)
	default:
		response := response.Response{
			Status:  "Unauthorized",
			Code:    fiber.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		}

		return c.Status(401).JSON(response)
	}
}
