package models

import (
	"fmt"

	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

// type Product struct {
// 	Id            int            `db:"id" json:"id"`
// 	Name          string         `db:"name" json:"name"`
// 	Description   sql.NullString `db:"description" json:"description"`
// 	Image         sql.NullString `db:"image" json:"image"`
// 	Discount      sql.NullInt64  `db:"discount" json:"discount"`
// 	IsRecommended bool           `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
// 	BasePrice     int            `db:"basePrice" json:"basePrice"`
// 	Category      sql.NullString `db:"category" json:"category"`
// 	Tag           sql.NullString `db:"tag" json:"tag"`
// 	Rating        sql.NullInt64  `db:"rating" json:"rating"`
// 	CreatedAt     time.Time      `db:"createdAt" json:"createdAt"`
// 	UpdatedAt     sql.NullTime   `db:"updatedAt" json:"updatedAt"`
// }

// type Product struct {
// 	Id            int            `db:"id" json:"id"`
// 	Name          string         `db:"name" json:"name" form:"name"`
// 	BasePrice     *int           `db:"basePrice" json:"basePrice" form:"basePrice"`
// 	Description   *string        `db:"description" json:"description" form:"description"`
// 	Image         sql.NullString `db:"image" json:"image"`
// 	IsRecommended *bool          `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
// 	Discount      *float64       `db:"discount" json:"discount" form:"discount"`
// 	CreatedAt     *time.Time     `db:"createdAt" json:"createdAt"`
// 	UpdatedAt     *time.Time     `db:"updatedAt" json:"updatedAt"`
// }

// func FindAllProduct(searchKey string, category string, sortBy string, order string, limit int, offset int) (services.InfoProduct, error) {
func FindAllProduct(searchKey string, sortBy string, order string, limit int, offset int) (services.InfoProduct, error) {

	sql := `
	SELECT * 
	FROM "products"
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	// sql := `SELECT
	// "p"."id",
	// "p"."name",
	// "p"."description",
	// "p"."basePrice",
	// "p"."image",
	// "p"."discount",
	// "c"."name" AS "category",
	// "t"."name" as "tag",
	// FROM "products" "p"
	// LEFT JOIN "productRatings" "pr" ON ("pr"."productId" = "p"."id")
	// LEFT JOIN "productCategories" "pc" on ("pc"."productId" = "p"."id")
	// LEFT JOIN "productTags" "pt" on ("pt"."productId" = "p"."id")
	// LEFT JOIN "categories" "c" on ("c"."id" = "pc"."categoryId")
	// LEFT join "tags" "t" on ("t"."id" = "pt"."tagId")
	// WHERE "p"."name" ILIKE $1 AND "c"."name" ILIKE $2
	// GROUP BY "p"."id", "c"."name", "t"."name"
	// ORDER BY "p"."` + sortBy + `" ` + order + `
	// LIMIT $3 OFFSET $4`
	fmt.Println(sql)

	sqlCount := `
	SELECT COUNT(*)
    FROM "products"
	WHERE "name" ILIKE $1
	`
	// sqlCount := `
	// SELECT COUNT(*) FROM "products" "p"
	// LEFT JOIN "productCategories" "pc" on ("pc"."productId" = "p"."id")
	// LEFT JOIN "productTags" "pt" on ("pt"."productId" = "p"."id")
	// LEFT JOIN "categories" "c" on ("c"."id" = "pc"."categoryId")
	// WHERE "p"."name" ILIKE $1 AND "c"."name" ILIKE $2
	// `
	fmt.Println(sqlCount)

	result := services.InfoProduct{}
	dataProduct := []services.Product{}

	err := db.Select(&dataProduct, sql, "%"+searchKey+"%", limit, offset)

	result.Data = dataProduct

	// row := db.QueryRow(sqlCount, "%"+searchKey+"%", "%"+category+"%")
	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)
	fmt.Println(result.Data)
	return result, err

}

// func FindAllProduct(searchKey string, limit int, offset int) (services.InfoProduct, error) {

// 	sql := `SELECT * FROM "products" WHERE "name" ILIKE $1 ORDER BY "id" ASC Limit $2 OFFSET $3`

// 	sqlCount := `SELECT COUNT(*) FROM "products" WHERE "name" ILIKE $1`
// 	result := services.InfoProduct{}
// 	dataProduct := []Product{}

// 	err := db.Select(&dataProduct, sql, "%"+searchKey+"%", limit, offset)

// 	result.Data = dataProduct

// 	row := db.QueryRow(sqlCount, "%"+searchKey+"%")

// 	err = row.Scan(&result.Count)
// 	return result, err
// }

func FindOneProduct(id int) (services.Product, error) {
	sql := `SELECT * FROM "products" WHERE "id" = $1`
	data := services.Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneProductByName(name string) (services.Product, error) {
	sql := `SELECT * FROM "products" WHERE "name" = $1`
	data := services.Product{}
	err := db.Get(&data, sql, name)
	return data, err
}

func CreateProduct(data services.ProductForm) (services.ProductForm, error) {
	sql := `
	INSERT INTO "products" ("name", "basePrice", "description", "image", "discount", "isRecommended")
	VALUES (:name, :basePrice, :description, COALESCE(NULLIF(:image,''),'default.png')) :discount,  :isRecommended)
	RETURNING *
	`
	result := services.ProductForm{}
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

func UpdateProduct(data services.ProductForm) (services.ProductForm, error) {
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
	result := services.ProductForm{}
	rows, err := db.NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteProduct(id int) (services.ProductForm, error) {
	sql := `DELETE FROM "products" WHERE "id" = $1 RETURNING *`
	data := services.ProductForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
