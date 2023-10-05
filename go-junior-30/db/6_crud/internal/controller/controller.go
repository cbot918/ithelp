package controller

import (
	db "6/db/sqlc"
	"6/internal/service"
)

type Controller struct {
	S service.Service
}

func NewController(q *db.Queries) *Controller {
	return &Controller{S: *service.NewService(q)}
}
