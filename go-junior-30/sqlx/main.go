package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// define User struct
type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func main() {

	var err error

	// load .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// connect db
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ping db
	err = db.Ping()
	if err != nil {
		fmt.Println("ping failed")
		return
	}
	fmt.Println("ping success")

	// Insert a user
	newUser := User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
	}
	insertUserSQL := "INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id"
	_, err = db.NamedExec(insertUserSQL, newUser)
	if err != nil {
		panic(err)
	}

	// select a user
	u := []User{}
	q := "SELECT * FROM users ORDER BY id DESC LIMIT 1"
	err = db.Select(&u, q)
	if err != nil {
		log.Fatal(err)
	}
	printJSON(u)

}

func printJSON(v any) {
	json, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(json))
}
