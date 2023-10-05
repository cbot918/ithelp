package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	db "6/db/sqlc"

	"6/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := internal.NewConfig()
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

	api := internal.NewAPI(app, query)

	err = api.APP.Listen(cfg.HOST)
	if err != nil {
		log.Fatal(err)
	}

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
