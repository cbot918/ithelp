package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())    // cors allow *
	app.Use(logger.New())  // logger middleware
	app.Use(recover.New()) // recover middleware

	app.Static("/", "public")

	app.Get("/hello", Hello)

	app.Post("/login", Login)

	if err := app.Listen(":8887"); err != nil {
		log.Fatal(err)
	}
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

type userParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	user := &userParam{}
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(422).SendString("data error")
	}

	fmt.Printf("%+v", user)

	return c.JSON(user)
}
