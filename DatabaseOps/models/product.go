package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // go get github.com/lib/pq ile modülü eklemek gerekebilir
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "AdventureWorks"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Price       float32
}

var db *sql.DB

func init() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disabled", host, port, user, password, dbname)
	// sqlDriver'ı bulmak için https://go.dev/wiki/SQLDrivers
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO product(title,description,price) VALUES($1,$2,$3)", data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}

	rowAffacted, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Affected row count is ", rowAffacted)
}
