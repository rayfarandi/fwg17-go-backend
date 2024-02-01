package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id        int          `db:"id" json:"id"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	Name      string       `db:"name" json:"name"`
	CreatedAt time.Time    `db:"createdAt" json:"createdat"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedat"`
}

func FindAllUsers() ([]User, error) {
	sql := `SELECT * FROM "users"`
	data := []User{}
	err := db.Select(&data, sql)
	return data, err
}
