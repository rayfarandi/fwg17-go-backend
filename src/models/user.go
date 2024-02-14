package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

var db *sqlx.DB = lib.DB

func FindAllUsers(searchKey string, sortBy string, order string, limit int, offset int) (services.InfoUser, error) {

	sql := `SELECT * FROM "users" WHERE "fullName" ILIKE $1 ORDER BY "` + sortBy + `"` + order + ` LIMIT $2 OFFSET $3`

	sqlCount := `SELECT COUNT(*) FROM "users" WHERE "fullName" ILIKE $1`
	result := services.InfoUser{}
	dataUser := []services.User{}

	err := db.Select(&dataUser, sql, "%"+searchKey+"%", limit, offset)
	result.Data = dataUser

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)
	return result, err
}

func FindOneUser(id int) (services.User, error) {
	sql := `SELECT * FROM "users" WHERE "id" = $1`
	data := services.User{}
	err := db.Get(&data, sql, id)
	return data, err
}
func FindOneUserEmail(email string) (services.User, error) {
	sql := `SELECT * FROM "users" WHERE "email" = $1`
	data := services.User{}
	err := db.Get(&data, sql, email)
	return data, err
}
func FindOneUserByEmail(email string) (services.User, error) {
	sql := `SELECT * FROM "users" WHERE "email" = $1`
	data := services.User{}
	err := db.Get(&data, sql, email)
	return data, err
}

func CreateUser(data services.UserForm) (services.UserForm, error) {
	sql := `
	INSERT INTO "users" ("email", "password", "fullName", "phoneNumber", "address", "role", "picture")
	VALUES (:email, :password, :fullName, :phoneNumber, :address, COALESCE(NULLIF(:role,''),'customer'), COALESCE(NULLIF(:picture,''),'default.png'))
	RETURNING *
	`
	result := services.UserForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	if rows.Next() {
		if err := rows.StructScan(&result); err != nil {
			return result, err
		}
	}

	return result, nil
}

func UpdateUser(data services.UserForm) (services.UserForm, error) {
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
	result := services.UserForm{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteUser(id int) (services.UserForm, error) {
	sql := `DELETE FROM "users" WHERE "id" = $1 RETURNING *`
	data := services.UserForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
