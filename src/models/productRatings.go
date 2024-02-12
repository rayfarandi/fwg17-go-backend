package models

import (
	"time"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

type ProductRatings struct {
	Id            int        `db:"id" json:"id"`
	ProductId     *int       `db:"productId" json:"productId" form:"productId"`
	Rate          *int       `db:"rate" json:"rate" form:"rate"`
	ReviewMessage *string    `db:"reviewMessage" json:"reviewMessage" form:"reviewMessage"`
	UserId        *int       `db:"userId" json:"userId" form:"userId"`
	CreatedAt     *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updatedAt" json:"updatedAt"`
}

func FindAllProductRatings(searchKey string, limit int, offset int) (services.InfoProductRatings, error) {

	sql := `SELECT * FROM "productRatings" WHERE "reviewMessage" ILIKE $1 ORDER BY "id" ASC Limit $2 OFFSET $3`

	sqlCount := `SELECT COUNT(*) FROM "productRatings" WHERE "reviewMessage" ILIKE $1`
	result := services.InfoProductRatings{}
	dataProductRatings := []ProductRatings{}

	err := db.Select(&dataProductRatings, sql, "%"+searchKey+"%", limit, offset)

	result.Data = dataProductRatings

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")

	err = row.Scan(&result.Count)
	return result, err
}

func FindOneProductRatings(id int) (ProductRatings, error) {
	sql := `SELECT * FROM "productRatings" WHERE "id" = $1`
	data := ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductRatingsByName(name string) (ProductRatings, error) {
	sql := `SELECT * FROM "productRatings" WHERE "reviewMessage" = $1`
	data := ProductRatings{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateProductRatings(data ProductRatings) (ProductRatings, error) {
	sql := `
	INSERT INTO "productRatings" ("productId", "rate", "reviewMessage", "userId")
	VALUES (:productId, :rate, :reviewMessage, :userId)
	RETURNING *
	`
	result := ProductRatings{}
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

func UpdateProductRatings(data ProductRatings) (ProductRatings, error) {
	sql := `UPDATE "productRatings" SET 
	"productId"=COALESCE(NULLIF(:productId,0),"productId"),
	"rate"=COALESCE(NULLIF(:rate,0),"rate"),
	"reviewMessage"=COALESCE(NULLIF(:reviewMessage,''),"reviewMessage"),
	"userId"=COALESCE(NULLIF(:userId,0),"userId"),
	"updatedAt"=NOW()
	WHERE "id"=:id
	RETURNING *
	`
	result := ProductRatings{}

	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductRatings(id int) (ProductRatings, error) {
	sql := `DELETE FROM "productRatings" WHERE "id" = $1 RETURNING *`
	data := ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}
