package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
	dsn    = "postgres://postgres:12345@localhost:5435/blog?sslmode=disable"
)

func main() {

	conn, err := newConn(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	s := NewStorage(conn)

	u := &User{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "12345",
		Discount: 0.5,
	}

	result, err := s.CreateUser(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	users, err := s.ReadUser()
	if err != nil {
		log.Fatal(err)
	}

	json, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))

}

// create connection
func newConn(driver string, dsn string) (*sql.DB, error) {
	conn, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("postgres good")
	return conn, nil
}

// types
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Discount float32
}

// CRUD functions
type Storage struct {
	Conn *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {
	return &Storage{Conn: conn}
}

func (s *Storage) CreateUser(u *User) (sql.Result, error) {

	query := `INSERT INTO users(name, email, password, discount) VALUES ($1, $2, $3, $4);`

	result, err := s.Conn.Exec(query, u.Name, u.Email, u.Password, u.Discount)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Storage) ReadUser() ([]*User, error) {

	query := `select * from users`

	rows, err := s.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Discount)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}
