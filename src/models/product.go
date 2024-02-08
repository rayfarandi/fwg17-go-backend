package models

import (
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type Product struct {
	Id            int        `db:"id" json:"id"`
	Name          *string    `db:"name" json:"name" form:"name"`
	BasePrice     *int       `db:"basePrice" json:"basePrice" form:"basePrice"`
	Description   *string    `db:"description" json:"description" form:"description"`
	Image         *string    `db:"image" json:"image" form:"image"`
	IsRecommended *bool      `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	Discount      *float64   `db:"discount" json:"discount" form:"discount"`
	CreatedAt     *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProduct(searchKey string, limit int, offset int) (services.InfoProduct, error) {

	sql := `SELECT * FROM "products" WHERE "name" ILIKE $1 ORDER BY "id" ASC Limit $2 OFFSET $3`

	sqlCount := `SELECT COUNT(*) FROM "products" WHERE "name" ILIKE $1`
	result := services.InfoProduct{}
	dataProduct := []Product{}

	err := db.Select(&dataProduct, sql, "%"+searchKey+"%", limit, offset)

	result.Data = dataProduct

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")

	err = row.Scan(&result.Count)
	return result, err
}

func FindOneProduct(id int) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "id" = $1`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductByName(name string) (Product, error) {
	sql := `SELECT * FROM "products" WHERE "name" = $1`
	data := Product{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateProduct(data Product) (Product, error) {
	sql := `
	INSERT INTO "product" ("name", "basePrice", "description", "images", "discount", "isRecommended")
	VALUES (:email, :basePrice, :description, :images, :discount,  :isRecommended)
	RETURNING *
	`
	result := Product{}
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

func UpdateProduct(data Product) (Product, error) {
	sql := `UPDATE "products" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"basePrice"=COALESCE(NULLIF(:basePrice,0),"basePrice"),
	"description"=COALESCE(NULLIF(:description,''),"description"),
	"image"=COALESCE(NULLIF(:image,''),"image"),
	"isRecommended"=COALESCE(:isRecommended,false),
	"discount"=COALESCE(NULLIF(:discount,0.0),"discount"),
	"updatedAt"=NOW()
	WHERE "id"=:id
	RETURNING *
	`
	result := Product{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProduct(id int) (Product, error) {
	sql := `DELETE FROM "products" WHERE "id" = $1 RETURNING *`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}
