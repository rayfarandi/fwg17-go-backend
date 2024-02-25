package lib

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func connectDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=go-coffee-go password=1 sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return db
}

var DB *sqlx.DB = connectDB()
