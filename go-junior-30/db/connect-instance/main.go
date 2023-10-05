package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// docker command
// docker run -it --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=test postgres

const dsn = "postgresql://postgres:12345@localhost:5432/test?sslmode=disable"

func main() {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("ping db success!")

	res, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, email text, password text)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
