package lib

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func connectDB() *sqlx.DB {
	// db, err := sqlx.Connect("postgres", "user=postgres dbname=go-coffee-go password=1 sslmode=disable") // local
	// db, err := sqlx.Connect("postgres", "user=postgres dbname=go-coffee-go password=1 sslmode=disable") // try docker compose

	dbConnect := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	db, err := sqlx.Connect("postgres", dbConnect) //Supabase

	if err != nil {
		fmt.Println(err)
	}
	return db
}

var DB *sqlx.DB = connectDB()
