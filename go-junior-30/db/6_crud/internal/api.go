package internal

import (
	db "6/db/sqlc"
	"6/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	C controller.Controller
}

func NewAPI(q *db.Queries) *API {
	return &API{C: *controller.NewController(q)}
}

func (a *API) Ping(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"message": "pong"})

}
