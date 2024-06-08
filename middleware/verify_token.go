package middleware

import (
	"go-todo-list/exception"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)

			claims := user.Claims.(jwt.MapClaims)

			c.Locals("claims", claims)

			return c.Next()
		},
		ErrorHandler: exception.CustomTokenErrorHandler,
	})
}
