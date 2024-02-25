package models

import (
	"fmt"

	"github.com/rayfarandi/fwg17-go-backend/src/service"
)

func FindAllCategories(searchKey string, sortBy string, order string, limit int, offset int) (service.InfoC, error) {
	sql := `
	SELECT * FROM "categories" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "categories"
	WHERE "name" ILIKE $1
	`

	result := service.InfoC{}
	data := []service.Categories{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneCategories(id int) (service.Categories, error) {
	sql := `SELECT * FROM "categories" WHERE id = $1`
	data := service.Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateCategories(data service.Categories) (service.Categories, error) {
	sql := `INSERT INTO "categories" ("name") VALUES (:name) RETURNING *`
	result := service.Categories{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateCategories(data service.Categories) (service.Categories, error) {
	sql := `UPDATE "categories" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.Categories{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(err)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteCategories(id int) (service.Categories, error) {
	sql := `DELETE FROM "categories" WHERE id = $1 RETURNING *`
	data := service.Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}
