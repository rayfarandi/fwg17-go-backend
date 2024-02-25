package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllTags(searchKey string, sortBy string, order string, limit int, offset int) (service.InfoT, error) {
	sql := `
	SELECT * FROM "tags" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "tags"
	WHERE "name" ILIKE $1
	`

	result := service.InfoT{}
	data := []service.Tags{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneTags(id int) (service.Tags, error) {
	sql := `SELECT * FROM "tags" WHERE id = $1`
	data := service.Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateTags(data service.Tags) (service.Tags, error) {
	sql := `INSERT INTO "tags" ("name") VALUES (:name)
	RETURNING *
	`
	result := service.Tags{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateTags(data service.Tags) (service.Tags, error) {
	sql := `UPDATE "tags" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.Tags{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteTags(id int) (service.Tags, error) {
	sql := `DELETE FROM "tags" WHERE id = $1 RETURNING *`
	data := service.Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}
