package postgres

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
)

func (p *postgresRepo) CreateDriver(req models.Driver) (models.Driver, error) {
	var (
		res models.Driver
	)
	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}
	err = p.Db.Db.QueryRow("insert into drivers (id,phone) values ($1,$2) returning phone,id", uuid, req.Phone).Scan(&res.Phone, &res.Id)
	if err != nil {
		fmt.Println("Error while inserting", err)
		return res, err
	}
	return res, nil
}
