package models

import "github.com/rayfarandi/fwg17-go-backend/src/service"

func FindAllMessage(sortBy string, order string, limit int, offset int) (service.InfoM, error) {
	sql := `
	SELECT * FROM "message" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "message"`

	result := service.InfoM{}
	data := []service.Message{}
	err := db.Select(&data, sql, limit, offset)
	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneMessage(id int) (service.Message, error) {
	sql := `SELECT * FROM "message" WHERE id = $1`
	data := service.Message{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateMessage(data service.Message) (service.Message, error) {
	sql := `INSERT INTO "message" ("recipientId", "senderId", "text") 
	VALUES
	(:recipientId, :senderId, :text)
	RETURNING *
	`
	result := service.Message{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateMessage(data service.Message) (service.Message, error) {
	sql := `UPDATE "message" SET
	"recipientId"=COALESCE(NULLIF(:recipientId, 0),"recipientId"),
	"senderId"=COALESCE(NULLIF(:senderId, 0),"senderId"),
	"text"=COALESCE(NULLIF(:text, ''),"text"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := service.Message{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteMessage(id int) (service.Message, error) {
	sql := `DELETE FROM "message" WHERE id = $1 RETURNING *`
	data := service.Message{}
	err := db.Get(&data, sql, id)
	return data, err
}
