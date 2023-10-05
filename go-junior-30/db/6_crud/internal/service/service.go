package service

import (
	db "6/db/sqlc"
	"6/internal/repository"
)

type Service struct {
	R repository.Repository
}

func NewService(q *db.Queries) *Service {
	return &Service{R: *repository.NewRepository(q)}
}
