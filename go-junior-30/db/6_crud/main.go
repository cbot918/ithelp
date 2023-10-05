package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	db "6/db/sqlc"
	i "6/internal"

	"6/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := i.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := newConn(cfg.DB_DRIVER, cfg.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	query := db.New(conn)

	api := internal.NewAPI(query)

	app.Get("/ping", api.Ping)

	err = app.Listen(cfg.HOST)
	if err != nil {
		log.Fatal(err)
	}

	// user := db.CreateUserParams{
	// 	Name:     "sqlctest",
	// 	Email:    "sqlctest@gmail.com",
	// 	Password: "12345",
	// 	Discount: 0.7,
	// }

	// q := db.New(conn)

	// result, err := q.CreateUser(context.Background(), user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// printJson(result)

	// u, err := q.GetUser(context.Background(), 12)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// printJson(u)

	// err = q.DeleteUser(context.Background(), 12)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// users, err := q.ListUsers(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// printJson(users)

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

// 封個 printJson 好了
func printJson(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json marshal failed")
	}
	fmt.Println(string(json))
}
