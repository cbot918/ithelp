package repository

import db "6/db/sqlc"

type Repository struct {
	Q *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{Q: q}
}
