package models

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"
)

type Product struct {
	Id            int            `db:"id" json:"id"`
	Name          string         `db:"name" json:"name"`
	Description   sql.NullString `db:"description" json:"description"`
	Image         sql.NullString `db:"image" json:"image"`
	Discount      sql.NullInt64  `db:"discount" json:"discount"`
	IsRecommended bool           `db:"isRecommended" json:"isRecommended"`
	TagId         int            `db:"tagId" json:"tagId"`
	BasePrice     int            `db:"basePrice" json:"basePrice"`
	Category      sql.NullString `db:"category" json:"category"`
	Tag           sql.NullString `db:"tag" json:"tag"`
	Rating        sql.NullInt64  `db:"rating" json:"rating"`
	CreatedAt     time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type ProductDetails struct {
	Id               int            `db:"id" json:"id"`
	Name             string         `db:"name" json:"name"`
	Description      sql.NullString `db:"description" json:"description"`
	BasePrice        int            `db:"basePrice" json:"basePrice"`
	Image            string         `db:"image" json:"image"`
	Discount         sql.NullInt64  `db:"discount" json:"discount"`
	IsRecommended    sql.NullBool   `db:"isRecommended" json:"isRecommended"`
	Tag              sql.NullString `db:"tag" json:"tag"`
	Rating           sql.NullInt64  `db:"rating" json:"rating"`
	Review           sql.NullInt64  `db:"review" json:"review"`
	VariantsProducts string         `db:"variantsProducts" json:"variantsProducts"`
	ProductImages    string         `db:"productImages" json:"productImages"`
	CreatedAt        time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt        sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type ProductForm struct {
	Id            int          `db:"id" json:"id"`
	Name          *string      `db:"name" json:"name" form:"name"`
	Description   *string      `db:"description" json:"description"`
	BasePrice     *int         `db:"basePrice" json:"basePrice" form:"basePrice"`
	Image         string       `db:"image" json:"image" form:"image"`
	Discount      *int         `db:"discount" json:"discount" form:"discount"`
	IsRecommended *bool        `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	TagId         *int         `db:"tagId" json:"tagId" form:"tagId"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoP struct {
	Data  []Product
	Count int
}

func FindAllProducts(searchKey string, category string, sortBy string, order string, limit int, offset int) (InfoP, error) {
	sql := `
	SELECT
	"p"."id",
	"p"."name",
	"p"."description",
	"p"."basePrice",
	"p"."image",
	"p"."discount",
	"c"."name" AS "category",
	"t"."name" as "tag",
	sum("pr"."rate")/count("pr"."id") as "rating"
	FROM "products" "p"
	LEFT JOIN "productRatings" "pr" ON ("pr"."productId" = "p"."id")
	LEFT JOIN "productCategories" "pc" on ("pc"."productId" = "p"."id")
	LEFT JOIN "categories" "c" on ("c"."id" = "pc"."categoryId")
	LEFT join "tags" "t" on ("t"."id" = "p"."tagId")
	WHERE "p"."name" ILIKE $1 AND "c"."name" ILIKE $2
	GROUP BY "p"."id", "c"."name", "t"."name"
	ORDER BY "p"."` + sortBy + `" ` + order + `
	LIMIT $3 OFFSET $4
	`
	sqlCount := `
	SELECT COUNT(*) FROM "products" "p"
	LEFT JOIN "productCategories" "pc" on ("pc"."productId" = "p"."id")
	LEFT JOIN "categories" "c" on ("c"."id" = "pc"."categoryId")
	WHERE "p"."name" ILIKE $1 AND "c"."name" ILIKE $2
	`

	result := InfoP{}
	data := []Product{}
	err := db.Select(&data, sql, "%"+searchKey+"%", "%"+category+"%", limit, offset)
	if err != nil {
		return result, err
	}

	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%", "%"+category+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProducts(id int) (ProductDetails, error) {
	sql := `
    SELECT
    "p"."id",
    "p"."name",
    "p"."description",
    "p"."basePrice",
    "p"."image",
    "p"."discount",
    "p"."createdAt",
    "p"."isRecommended",
    "t"."name" as "tag",
    sum("pr"."rate")/count("pr"."id") AS "rating",
    count(DISTINCT "pr"."id") AS "review",
    JSONB_AGG(
        DISTINCT JSONB_BUILD_OBJECT(
            'id', "v"."id",
            'name', "v"."name",
            'additionalPrice', "v"."additionalPrice"
        )
    ) AS "variantsProducts",
	JSONB_AGG(
        DISTINCT JSONB_BUILD_OBJECT(
            'id', "pi"."id",
            'imageUrl', "pi"."imageUrl"
        )
    ) AS "productImages"
    FROM "products" "p"
    LEFT JOIN "productRatings" "pr" ON ("pr"."productId" = "p"."id")
    LEFT JOIN "tags" "t" on ("t"."id" = "p"."tagId")
    LEFT JOIN "productVariant" "pv" on ("pv"."productId" = "p"."id")
    LEFT JOIN "variant" "v" on ("pv"."variantId" = "v"."id")
	LEFT JOIN "productImages" "pi" on ("pi"."productId" = "p"."id")
    WHERE "p"."id" = $1
    GROUP BY "p"."id", "t"."name"
    `
	data := ProductDetails{}
	encoded := base64.StdEncoding.EncodeToString([]byte(data.VariantsProducts))
	decodedVariants, error := base64.StdEncoding.DecodeString(encoded)
	if error != nil {
		fmt.Println("decode error:", error)
		return data, error
	}

	data.VariantsProducts = string(decodedVariants)
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProducts(data ProductForm) (ProductForm, error) {
	sql := `INSERT INTO "products" ("name", "description", "basePrice", "image", "discount", "isRecommended", "tagId") 
	VALUES
	(:name, :description, :basePrice, :image, :discount, :isRecommended, :tagId)
	RETURNING *
	`
	result := ProductForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProduct(data ProductForm) (ProductForm, error) {
	sql := `UPDATE "products" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"description"=COALESCE(NULLIF(:description, ''),"description"),
	"basePrice"=COALESCE(NULLIF(:basePrice, 0),"basePrice"),
	"image"=COALESCE(NULLIF(:image, ''),"image"),
	"discount"=COALESCE(NULLIF(:discount, 0),"discount"),
	"isRecommended"=COALESCE(NULLIF(:isRecommended, false),"isRecommended"),
	"tagId"=COALESCE(NULLIF(:tagId, 0),"tagId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := ProductForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProduct(id int) (ProductForm, error) {
	sql := `DELETE FROM "products" WHERE id = $1 RETURNING *`
	data := ProductForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
