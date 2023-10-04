package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	port      = ":3001"
	staticDir = "./ui/dist"
)

func main() {
	app := fiber.New()

	app.Static("/", staticDir)

	fmt.Printf("Server is listening on port %s\n", port)

	err := app.Listen(port)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
