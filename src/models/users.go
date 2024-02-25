package models

import (
	"fmt"

	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/service"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

func FindAllUsers(searchKey string, sortBy string, order string, limit int, offset int) (service.Info, error) {
	sql := `
	SELECT * FROM "users" 
	WHERE "fullName" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "users"
	WHERE "fullName" ILIKE $1
	`

	result := service.Info{}
	data := []service.User{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneUsers(id int) (service.User, error) {
	sql := `SELECT * FROM "users" WHERE id = $1`
	data := service.User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneUsersByEmail(email string) (service.User, error) {
	sql := `SELECT * FROM "users" WHERE email = $1`
	data := service.User{}
	err := db.Get(&data, sql, email)
	return data, err
}

func CreateUser(data service.UserForm) (service.UserForm, error) {
	fmt.Println(data.PhoneNumber)
	fmt.Println(data.Picture)

	sql := `INSERT INTO "users" ("fullName", "email", "password", "address", "phoneNumber", "role", "picture") 
	VALUES
	(:fullName, :email, :password, :address, :phoneNumber, :role, :picture)
	RETURNING *
	`
	result := service.UserForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateUser(data service.UserForm) (service.UserForm, error) {
	sql := `UPDATE "users" SET
	"fullName"=COALESCE(NULLIF(:fullName, ''),"fullName"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"password"=COALESCE(NULLIF(:password, ''),"password"),
	"address"=COALESCE(NULLIF(:address, ''),"address"),
	"picture"=COALESCE(NULLIF(:picture, ''),"picture"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber, ''),"phoneNumber"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.UserForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteUser(id int) (service.UserForm, error) {
	sql := `DELETE FROM "users" WHERE id = $1 RETURNING *`
	data := service.UserForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
