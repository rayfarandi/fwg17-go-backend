package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllVariants(searchKey string, sortBy string, order string, limit int, offset int) (service.InfoV, error) {
	sql := `
	SELECT * FROM "variant" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "variant"
	WHERE "name" ILIKE $1
	`

	result := service.InfoV{}
	data := []service.Variants{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneVariants(id int) (service.Variants, error) {
	sql := `SELECT * FROM "variant" WHERE id = $1`
	data := service.Variants{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateVariants(data service.Variants) (service.Variants, error) {
	sql := `INSERT INTO "variant" ("name", "additionalPrice") VALUES (:name, :additionalPrice) RETURNING *`
	result := service.Variants{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateVariants(data service.Variants) (service.Variants, error) {
	sql := `UPDATE "variant" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, 0),"additionalPrice"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.Variants{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteVariants(id int) (service.Variants, error) {
	sql := `DELETE FROM "variant" WHERE id = $1 RETURNING *`
	data := service.Variants{}
	err := db.Get(&data, sql, id)
	return data, err
}
