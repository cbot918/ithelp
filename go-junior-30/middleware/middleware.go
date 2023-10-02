package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// auth middleware
func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		_, err := DecodeJWT(token)
		if err != nil {
			fmt.Println("422 in here")
			return c.Status(422).JSON(fiber.Map{"message": "token is invalid, please check your token"})
		}
		return c.Next()
	}
}
