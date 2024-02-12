package models

import (
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type ProductTags struct {
	Id        int        `db:"id" json:"id"`
	TagId     *int       `db:"tagId" json:"tagId" form:"tagId"`
	ProductId *int       `db:"productId" json:"productId" form:"productId"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductTags(limit int, offset int) (services.InfoProductTags, error) {

	sql := `SELECT * FROM "productTags" ORDER BY "id" ASC Limit $1 OFFSET $2`

	sqlCount := `SELECT COUNT(*) FROM "productTags"`
	result := services.InfoProductTags{}
	dataProductTags := []ProductTags{}

	err := db.Select(&dataProductTags, sql, limit, offset)

	result.Data = dataProductTags

	row := db.QueryRow(sqlCount)

	err = row.Scan(&result.Count)
	return result, err
}

func FindOneProductTags(id int) (ProductTags, error) {
	sql := `SELECT * FROM "productTags" WHERE "id" = $1`
	data := ProductTags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductTags(data ProductTags) (ProductTags, error) {
	sql := `
	INSERT INTO "productTags" ("tagId","productId")
	VALUES (:tagId, :productId)
	RETURNING *
	`
	result := ProductTags{}
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

func UpdateProductTags(data ProductTags) (ProductTags, error) {
	sql := `UPDATE "productTags" SET 
	"tagId"=COALESCE(NULLIF(:tagId,0),"tagId"),
	"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"updatedAt"=NOW()
	WHERE "id"=:id
	RETURNING *
	`
	result := ProductTags{}

	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductTags(id int) (ProductTags, error) {
	sql := `DELETE FROM "productTags" WHERE "id" = $1 RETURNING *`
	data := ProductTags{}
	err := db.Get(&data, sql, id)
	return data, err
}
