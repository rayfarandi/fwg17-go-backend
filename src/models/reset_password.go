package models

import (
	"database/sql"
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type FormReset struct {
	Id        int          `db:"id"`
	Email     string       `db:"email" form:"email"`
	Otp       string       `db:"otp" form:"otp"`
	CreatedAt time.Time    `db:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt"`
}

func FindOneRPByOtp(otp string) (FormReset, error) {
	sql := `SELECT * FROM "resetPassword" WHERE otp=$1`
	data := FormReset{}
	err := db.Get(&data, sql, otp)
	return data, err
}

func CreateResetPassword(data FormReset) (FormReset, error) {
	sql := `
	INSERT INTO "resetPassword" ("email","otp") VALUES (:email, :otp)
	RETURNING *`

	result := FormReset{}
	rows, err := db.NamedQuery(sql, data)
	// if err != nil {
	// 	return result, err
	// }

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteResetPassword(id int) (services.User, error) {
	sql := `DELETE FROM "resetPassword" WHERE "id" = $1 RETURNING *`
	data := services.User{}
	err := db.Get(&data, sql, id)
	return data, err
}

// type FormResetPassword struct {
// 	Id        int          `db:"id"`
// 	Email     string       `db:"email" form:"email"`
// 	Otp       string       `db:"otp" form:"otp"`
// 	CreatedAt time.Time    `db:"createdAt"`
// 	UpdatedAt sql.NullTime `db:"updatedAt"`
// }

// func FindOneRPByOtp(otp string) (FormResetPassword, error) {
// 	sql := `SELECT * FROM "resetPassword" WHERE "otp" = $1`
// 	data := FormResetPassword{}
// 	err := db.Get(&data, sql, otp)
// 	return data, err
// }

// func CreateResetPassword(data FormResetPassword) (FormResetPassword, error) {
// 	sql := `
// 	INSERT INTO "resetPassword" ("email","otp") VALUES
// 	(:email, :otp)
// 	RETURNING *`

// 	result := FormResetPassword{}
// 	rows, err := db.NamedQuery(sql, data)

// 	for rows.Next() {
// 		rows.StructScan(&result)
// 	}

// 	return result, err
// }

// func DeleteResetPassword(id int) (User, error) {
// 	sql := `DELETE FROM "resetPassword" WHERE "id" = $1 RETURNING *`
// 	data := User{}
// 	err := db.Get(&data, sql, id)
// 	return data, err
// }
