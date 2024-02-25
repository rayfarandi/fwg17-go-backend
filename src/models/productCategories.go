package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllProductCategories(sortBy string, order string, limit int, offset int) (service.InfoPC, error) {
	sql := `
	SELECT * FROM "productCategories" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "productCategories"
	`

	result := service.InfoPC{}
	data := []service.ProductCategories{}
	err := db.Select(&data, sql, limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductCategories(id int) (service.ProductCategories, error) {
	sql := `SELECT * FROM "productCategories" WHERE id = $1`
	data := service.ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductCategories(data service.ProductCategories) (service.ProductCategories, error) {
	sql := `INSERT INTO "productCategories" ("productId", "categoryId") 
	VALUES
	(:productId, :categoryId)
	RETURNING *
	`
	result := service.ProductCategories{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductCategories(data service.ProductCategories) (service.ProductCategories, error) {
	sql := `UPDATE "productCategories" SET
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"categoryId"=COALESCE(NULLIF(:categoryId, 0),"categoryId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.ProductCategories{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductCategories(id int) (service.ProductCategories, error) {
	sql := `DELETE FROM "productCategories" WHERE id = $1 RETURNING *`
	data := service.ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}
