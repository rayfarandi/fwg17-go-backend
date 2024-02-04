package models

import (
	"database/sql"
	"time"

	"github.com/LukaGiorgadze/gonull"
	"github.com/jmoiron/sqlx"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int                     `db:"id" json:"id"`
	Fullname    string                  `db:"fullName" json:"fullName" form:"fullName"`
	Email       string                  `db:"email" json:"email" form:"email"`
	Password    string                  `db:"password" json:"password" form:"password"`
	Address     gonull.Nullable[string] `db:"address" json:"address" form:"address"`
	Picture     gonull.Nullable[string] `db:"picture" json:"picture" form:"picture"`
	PhoneNumber gonull.Nullable[string] `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        string                  `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time               `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime            `db:"updatedAt" json:"updatedAt"`
}

type InfoUser struct {
	Data  []User
	Count int
}

func FindAllUsers() ([]User, error) {
	sql := `SELECT * FROM "users"`
	data := []User{}
	err := db.Select(&data, sql)
	return data, err
}

func FindOneUser(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE "id" = $1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateUser(data User) (User, error) {

	sql := `INSERT INTO "users" ("fullName", "email", "password", "role") VALUES (:fullName, :email, :password, :role) RETURNING *`
	result := User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func UpdateUser(data User) (User, error) {
	sql := `UPDATE "users" SET 
	"email"=COALESCE(NULLIF(:email,''),"email"),
	"password"=COALESCE(NULLIF(:password,''),"password"),
	"fullName"=COALESCE(NULLIF(:fullName,''),"fullName"),
	"address"=COALESCE(NULLIF(:address,''),"address"),
	"picture"=COALESCE(NULLIF(:picture,''),"picture"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber,''),"phoneNumber"),
	"role"=COALESCE(NULLIF(:role, ''),"role"),
	"updatedAt"=now()
	WHERE "id"=:id
	RETURNING *
	`
	result := User{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteUser(id int) (User, error) {
	sql := `DELETE FROM "users" WHERE "id" = $1 RETURNING *`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}
