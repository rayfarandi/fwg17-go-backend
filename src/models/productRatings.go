package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllProductRatings(sortBy string, order string, limit int, offset int) (service.InfoPR, error) {
	sql := `
	SELECT * FROM "productRatings" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "productRatings"
	`

	result := service.InfoPR{}
	data := []service.ProductRatings{}
	err := db.Select(&data, sql, limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductRatings(id int) (service.ProductRatings, error) {
	sql := `SELECT * FROM "productRatings" WHERE id = $1`
	data := service.ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductRatings(data service.PRForm) (service.PRForm, error) {
	sql := `INSERT INTO "productRatings" ("productId", "rate", "reviewMessage", "userId") 
	VALUES
	(:productId, :rate, :reviewMessage, :userId)
	RETURNING *
	`
	result := service.PRForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductRatings(data service.PRForm) (service.PRForm, error) {
	sql := `UPDATE "productRatings" SET
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"rate"=COALESCE(NULLIF(:rate, 0),"rate"),
	"reviewMessage"=COALESCE(NULLIF(:reviewMessage, ''),"reviewMessage"),
	"userId"=COALESCE(NULLIF(:userId, 0),"userId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.PRForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductRatings(id int) (service.PRForm, error) {
	sql := `DELETE FROM "productRatings" WHERE id = $1 RETURNING *`
	data := service.PRForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
