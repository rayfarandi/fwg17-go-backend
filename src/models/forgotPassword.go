package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllForgotPassword(searchKey string, sortBy string, order string, limit int, offset int) (service.InfoFP, error) {
	sql := `
	SELECT * FROM "forgotPassword" 
	WHERE "email" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "forgotPassword"
	WHERE "email" ILIKE $1
	`

	result := service.InfoFP{}
	data := []service.ForgotPassword{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneForgotPassword(id int) (service.ForgotPassword, error) {
	sql := `SELECT * FROM "forgotPassword" WHERE id = $1`
	data := service.ForgotPassword{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneByOtp(otp string) (service.ForgotPassword, error) {
	sql := `SELECT * FROM "forgotPassword" WHERE otp = $1`
	data := service.ForgotPassword{}
	err := db.Get(&data, sql, otp)
	return data, err
}

func CreateForgotPassword(data service.ForgotPassword) (service.ForgotPassword, error) {
	sql := `INSERT INTO "forgotPassword" ("otp", "email", "userId") 
	VALUES (:otp, :email, :userId) 
	RETURNING *
	`
	result := service.ForgotPassword{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateForgotPassword(data service.FPForm) (service.FPForm, error) {
	sql := `UPDATE "forgotPassword" SET
	"otp"=COALESCE(NULLIF(:otp, 0),"otp"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"updatedAt"=CURRENT_TIMESTAMP
	WHERE id=:id
	RETURNING *
	`
	result := service.FPForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteForgotPassword(id int) (service.FPForm, error) {
	sql := `DELETE FROM "forgotPassword" WHERE id = $1 RETURNING *`
	data := service.FPForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
