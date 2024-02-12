package models

import (
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type ProductCategories struct {
	Id         int        `db:"id" json:"id"`
	CategoryId *int       `db:"categoryId" json:"categoryId" form:"categoryId"`
	ProductId  *int       `db:"productId" json:"productId" form:"productId"`
	CreatedAt  *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductCategories(limit int, offset int) (services.InfoProductCategories, error) {

	sql := `SELECT * FROM "productCategories" ORDER BY "id" ASC Limit $1 OFFSET $2`

	sqlCount := `SELECT COUNT(*) FROM "productCategories"`
	result := services.InfoProductCategories{}
	dataProductCategories := []ProductCategories{}

	err := db.Select(&dataProductCategories, sql, limit, offset)

	result.Data = dataProductCategories

	row := db.QueryRow(sqlCount)

	err = row.Scan(&result.Count)
	return result, err
}

func FindOneProductCategories(id int) (ProductCategories, error) {
	sql := `SELECT * FROM "productCategories" WHERE "id" = $1`
	data := ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductCategories(data ProductCategories) (ProductCategories, error) {
	sql := `
	INSERT INTO "productCategories" ("categoryId","productId")
	VALUES (:categoryId, :productId)
	RETURNING *
	`
	result := ProductCategories{}
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

func UpdateProductCategories(data ProductCategories) (ProductCategories, error) {
	sql := `UPDATE "productCategories" SET 
	"categoryId"=COALESCE(NULLIF(:categoryId,0),"categoryId"),
	"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"updatedAt"=NOW()
	WHERE "id"=:id
	RETURNING *
	`
	result := ProductCategories{}

	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductCategories(id int) (ProductCategories, error) {
	sql := `DELETE FROM "productCategories" WHERE "id" = $1 RETURNING *`
	data := ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}
