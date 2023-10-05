package controller

import (
	db "6/db/sqlc"
	"6/internal/service"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	S service.Service
}

func NewController(q *db.Queries) *Controller {
	return &Controller{S: *service.NewService(q)}
}

func (c *Controller) Ping(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"message": "pong"})

}

func (c *Controller) CreateUser(ctx *fiber.Ctx) error {

	userReq := db.CreateUserParams{}

	err := ctx.BodyParser(&userReq)
	if err != nil {
		return ctx.Status(422).SendString("request body error")
	}

	fmt.Printf("%#+v", userReq)

	resUser, err := c.S.CreateUser(userReq)
	if err != nil {
		return err
	}
	return ctx.JSON(resUser)
}

func (c *Controller) GetUser(ctx *fiber.Ctx) error {
	idStr := ctx.Query("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return err
	}
	id32 := int32(id64)

	retUser, err := c.S.GetUser(id32)
	if err != nil {
		return err
	}

	return ctx.JSON(retUser)
}

func (c *Controller) DeleteUser(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return err
	}
	id32 := int32(id64)

	err = c.S.DeleteUser(id32)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "delete user success"})
}
