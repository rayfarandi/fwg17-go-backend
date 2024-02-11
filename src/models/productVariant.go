package models

import (
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type ProductVariant struct {
	Id              int        `db:"id" json:"id"`
	Name            *string    `db:"name" json:"name" form:"name"`
	ProductId       *int       `db:"productId" json:"productId" form:"productId"`
	AdditionalPrice *int       `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductVariant(searchKey string, limit int, offset int) (services.InfoProduct, error) {

	sql := `SELECT * FROM "productVariant" WHERE "name" ILIKE $1 ORDER BY "id" ASC Limit $2 OFFSET $3`

	sqlCount := `SELECT COUNT(*) FROM "productVariant" WHERE "name" ILIKE $1`
	result := services.InfoProduct{}
	dataProductVariant := []ProductVariant{}

	err := db.Select(&dataProductVariant, sql, "%"+searchKey+"%", limit, offset)

	result.Data = dataProductVariant

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")

	err = row.Scan(&result.Count)
	return result, err
}

func FindOneProductVariant(id int) (ProductVariant, error) {
	sql := `SELECT * FROM "productVariant" WHERE "id" = $1`
	data := ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductVariantByName(name string) (ProductVariant, error) {
	sql := `SELECT * FROM "productVariant" WHERE "name" = $1`
	data := ProductVariant{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateProductVariant(data ProductVariant) (ProductVariant, error) {
	sql := `
	INSERT INTO "productVariant" ("name", "productId", "additionalPrice")
	VALUES (:name, :productId, :additionalPrice)
	RETURNING *
	`
	result := ProductVariant{}
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

func UpdateProductVariant(data ProductVariant) (ProductVariant, error) {
	sql := `UPDATE "productVariant" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice,0),"additionalPrice"),
	"updatedAt"=NOW()
	WHERE "id"=:id
	RETURNING *
	`
	result := ProductVariant{}

	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductVariant(id int) (ProductVariant, error) {
	sql := `DELETE FROM "productVariant" WHERE "id" = $1 RETURNING *`
	data := ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}
