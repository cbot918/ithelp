package internal

import (
	db "6/db/sqlc"
	"6/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	APP *fiber.App
	C   controller.Controller
}

func NewAPI(app *fiber.App, q *db.Queries) *API {

	a := new(API)

	a.C = *controller.NewController(q)
	a.APP = app

	a.APP.Get("/ping", a.C.Ping)

	a.APP.Get("/getuser", a.C.GetUser)
	a.APP.Post("/createuser", a.C.CreateUser)
	a.APP.Delete("/deleteuser/:id", a.C.DeleteUser)
	return a
}
