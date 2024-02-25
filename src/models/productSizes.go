package models

import (
	"fmt"

	"github.com/rayfarandi/fwg17-go-backend/src/service"
)

func FindAllSizes(searchKey string, sortBy string, order string, limit int, offset int) (service.InfoS, error) {
	sql := `
	SELECT * FROM "sizes" 
	WHERE "size" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "sizes"
	WHERE "size" ILIKE $1
	`

	result := service.InfoS{}
	data := []service.Sizes{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	fmt.Println(result)
	return result, err
}

func FindOneSizes(id int) (service.Sizes, error) {
	sql := `SELECT * FROM "sizes" WHERE id = $1`
	data := service.Sizes{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateSizes(data service.SizesForm) (service.SizesForm, error) {
	sql := `INSERT INTO "sizes" ("size", "additionalPrice") 
	VALUES
	(:size, :additionalPrice)
	RETURNING *
	`
	result := service.SizesForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateSizes(data service.SizesForm) (service.SizesForm, error) {
	sql := `UPDATE "sizes" SET
	"size"=COALESCE(NULLIF(:size, ''),"size"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, 0),"additionalPrice"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.SizesForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteSizes(id int) (service.SizesForm, error) {
	sql := `DELETE FROM "sizes" WHERE id = $1 RETURNING *`
	data := service.SizesForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
