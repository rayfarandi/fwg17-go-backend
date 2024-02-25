package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllProductVariants(sortBy string, order string, limit int, offset int) (service.InfoPV, error) {
	sql := `
	SELECT * FROM "productVariant" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "productVariant"
	`

	result := service.InfoPV{}
	data := []service.ProductVariants{}
	err := db.Select(&data, sql, limit, offset)

	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductVariants(id int) (service.ProductVariants, error) {
	sql := `SELECT * FROM "productVariant" WHERE id = $1`
	data := service.ProductVariants{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductVariants(data service.ProductVariants) (service.ProductVariants, error) {
	sql := `INSERT INTO "productVariant" ("productId", "variantId") 
	VALUES
	(:productId, :variantId)
	RETURNING *
	`
	result := service.ProductVariants{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductVariants(data service.ProductVariants) (service.ProductVariants, error) {
	sql := `UPDATE "productVariant" SET
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"variantId"=COALESCE(NULLIF(:variantId, 0),"variantId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.ProductVariants{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductVariants(id int) (service.ProductVariants, error) {
	sql := `DELETE FROM "productVariant" WHERE id = $1 RETURNING *`
	data := service.ProductVariants{}
	err := db.Get(&data, sql, id)
	return data, err
}
