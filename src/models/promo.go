package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllPromo(searchKey string, sortBy string, order string, limit int, offset int) (service.InfoPo, error) {
	sql := `
	SELECT * FROM "promo" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "promo"
	WHERE "name" ILIKE $1
	`

	result := service.InfoPo{}
	data := []service.Promo{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOnePromo(id int) (service.Promo, error) {
	sql := `SELECT * FROM "promo" WHERE id = $1`
	data := service.Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreatePromo(data service.PromoForm) (service.PromoForm, error) {
	sql := `INSERT INTO "promo" ("name","code", "description", "percentage", "isExpired", "maximumPromo", "minimumAmount") 
	VALUES
	(:name, :code, :description, :percentage, :isExpired, :maximumPromo, :minimumAmount)
	RETURNING *
	`
	result := service.PromoForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdatePromo(data service.PromoForm) (service.PromoForm, error) {
	sql := `UPDATE "promo" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"code"=COALESCE(NULLIF(:code, ''),"code"),
	"description"=COALESCE(NULLIF(:description, ''),"description"),
	"percentage"=COALESCE(NULLIF(:percentage, 0),"percentage"),
	"isExpired"=COALESCE(NULLIF(:isExpired, false),"isExpired"),
	"maximumPromo"=COALESCE(NULLIF(:maximumPromo, 0),"maximumPromo"),
	"minimumAmount"=COALESCE(NULLIF(:minimumAmount, 0),"minimumAmount"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.PromoForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeletePromo(id int) (service.PromoForm, error) {
	sql := `DELETE FROM "promo" WHERE id = $1 RETURNING *`
	data := service.PromoForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
