package models

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rayfarandi/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int        `db:"id" json:"id"`
	Fullname    string     `db:"fullName" json:"fullName" form:"fullName"`
	Email       string     `db:"email" json:"email" form:"email"`
	Password    string     `db:"password" json:"password" form:"password"`
	Address     *string    `db:"address" json:"address" form:"address"`
	Picture     *string    `db:"picture" json:"picture" form:"picture"`
	PhoneNumber *string    `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        string     `db:"role" json:"role" form:"role"`
	CreatedAt   *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updatedAt" json:"updatedAt"`
}

type InfoUser struct {
	Data  []User
	Count int
}

func FindAllUsers(searchKey string, limit int, offset int) (InfoUser, error) {

	sql := `SELECT * FROM "users" WHERE "fullName" ILIKE $1 ORDER BY "id" ASC Limit $2 OFFSET $3`

	sqlCount := `SELECT COUNT(*) FROM "users" WHERE "fullName" ILIKE $1`
	result := InfoUser{}
	dataUser := []User{}

	err := db.Select(&dataUser, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		log.Printf("Error executing query: %v", err)
	}
	result.Data = dataUser

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")

	err = row.Scan(&result.Count)
	return result, err
}

// func FindAllUsers(limit int, offset int, searchKey string) (InfoUser, error) {
// 	sql := `SELECT * FROM "users" WHERE "fullName" ILIKE $3 ORDER BY "id" ASC LIMIT $1 OFFSET $2`
// 	sqlCount := `SELECT COUNT(*) FROM "users"`
// 	result := InfoUser{}
// 	dataUser := []User{}
// 	fmtSearch := fmt.Sprintf("%%%v%%", searchKey)
// 	err := db.Select(&dataUser, sql, limit, offset, fmtSearch)
// 	result.Data = dataUser
// 	row := db.QueryRow(sqlCount)
// 	err = row.Scan(&result.Count)
// 	// log.Println("error", sql)
// 	// log.Println("Parameters:", limit, offset, fmtSearch)

// 	return result, err
// }

func FindOneUser(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE "id" = $1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}
func FindOneUserEmail(email string) (User, error) {
	sql := `SELECT * FROM "users" WHERE "email" = $1`
	data := User{}
	err := db.Get(&data, sql, email)
	return data, err
}
func FindOneUserByEmail(email string) (User, error) {
	sql := `SELECT * FROM "users" WHERE "email" = $1`
	data := User{}
	err := db.Get(&data, sql, email)
	return data, err
}

// func CreateUser(data User) (User, error) {

// 	sql := `
// 	INSERT INTO "users" ("email", "password", "fullName", "phoneNumber", "address", "role", "picture")
// 	VALUES (:email, :password, :fullName, :phoneNumber, :address, COALESCE(:role, 'Customer'), :picture)
// 	RETURNING *
// 	`

// 	result := User{}
// 	rows, err := db.NamedQuery(sql, data)

// 	for rows.Next() {
// 		rows.StructScan(&result)
// 	}

//		return result, err
//	}
func CreateUser(data User) (User, error) {
	sql := `
	INSERT INTO "users" ("email", "password", "fullName", "phoneNumber", "address", "role", "picture")
	VALUES (:email, :password, :fullName, :phoneNumber, :address, COALESCE(NULLIF(:role,''),'customer'), :picture)
	RETURNING *
	`

	result := User{}
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
