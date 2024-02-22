package postgres

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
)

func (p *postgresRepo) CreateUser(req models.VerifyNumber) error {
	var (
		res string
	)
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	err = p.Db.Db.QueryRow("insert into users (id,phone) values ($1,$2) returning phone", uuid, req.Phone).Scan(&res)
	if err != nil {
		fmt.Println("Error while inserting", err)
		return err
	}
	return nil
}
