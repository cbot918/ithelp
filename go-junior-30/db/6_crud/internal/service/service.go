package service

import (
	db "6/db/sqlc"
	"context"
	"database/sql"
	"fmt"
)

type Service struct {
	Q *db.Queries
}

func NewService(q *db.Queries) *Service {
	return &Service{Q: q}
}

func (s *Service) CreateUser(user db.CreateUserParams) (*db.User, error) {

	resUser, err := s.Q.CreateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &resUser, err

}

func (s *Service) GetUser(id int32) (*db.User, error) {
	resUser, err := s.Q.GetUser(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &resUser, nil
}

func (s *Service) DeleteUser(id int32) error {

	_, err := s.Q.GetUser(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user id doesn't exists")
		}
		return err
	}

	err = s.Q.DeleteUser(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
