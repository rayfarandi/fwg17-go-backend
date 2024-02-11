package models

import (
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type Tags struct {
	Id        int        `db:"id" json:"id"`
	Name      *string    `db:"name" json:"name" form:"name"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllTags(searchKey string, limit int, offset int) (services.InfoTags, error) {

	sql := `SELECT * FROM "tags" WHERE "name" ILIKE $1 ORDER BY "id" ASC Limit $2 OFFSET $3`

	sqlCount := `SELECT COUNT(*) FROM "tags" WHERE "name" ILIKE $1`
	result := services.InfoTags{}
	dataTags := []Tags{}

	err := db.Select(&dataTags, sql, "%"+searchKey+"%", limit, offset)

	result.Data = dataTags

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")

	err = row.Scan(&result.Count)
	return result, err
}

func FindOneTags(id int) (Tags, error) {
	sql := `SELECT * FROM "tags" WHERE "id" = $1`
	data := Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneTagsByName(name string) (Tags, error) {
	sql := `SELECT * FROM "tags" WHERE "name" = $1`
	data := Tags{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateTags(data Tags) (Tags, error) {
	sql := `
	INSERT INTO "tags" ("name")
	VALUES (:name)
	RETURNING *
	`
	result := Tags{}
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

func UpdateTags(data Tags) (Tags, error) {
	sql := `UPDATE "tags" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"updatedAt"=NOW()
	WHERE "id"=:id
	RETURNING *
	`
	result := Tags{}

	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteTags(id int) (Tags, error) {
	sql := `DELETE FROM "tags" WHERE "id" = $1 RETURNING *`
	data := Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}
