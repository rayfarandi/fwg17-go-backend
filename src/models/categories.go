package models

import (
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type Categories struct {
	Id        int        `db:"id" json:"id"`
	Name      *string    `db:"name" json:"name" form:"name"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllCategories(searchKey string, limit int, offset int) (services.InfoCategories, error) {

	sql := `SELECT * FROM "categories" WHERE "name" ILIKE $1 ORDER BY "id" ASC Limit $2 OFFSET $3`

	sqlCount := `SELECT COUNT(*) FROM "categories" WHERE "name" ILIKE $1`
	result := services.InfoCategories{}
	data := []Categories{}

	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)

	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")

	err = row.Scan(&result.Count)
	return result, err
}

func FindOneCategories(id int) (Categories, error) {
	sql := `SELECT * FROM "categories" WHERE "id" = $1`
	data := Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateCategories(data Categories) (Categories, error) {
	sql := `
	INSERT INTO "categories" ("name") VALUES
	(:name)
	RETURNING *
	`
	result := Categories{}
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

func UpdateCategories(data Categories) (Categories, error) {
	sql := `UPDATE "categories" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"updatedAt"=NOW()
	WHERE id =:id
	RETURNING *
	`
	result := Categories{}

	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteCategories(id int) (Categories, error) {
	sql := `DELETE FROM "categories" WHERE "id" = $1 RETURNING *`
	data := Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}
