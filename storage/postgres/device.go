package postgres

import (
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
)

func (p *postgresRepo) CreateDeviceType(req models.DeviceType) (models.DeviceType, error) {
	var (
		res models.DeviceType
	)
	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}
	err = p.Db.Db.QueryRow("insert into device_type (id,name) values ($1,$2) returning id,name", uuid, req.Name).Scan(&res.Id, &res.Name)
	if err != nil {
		return res, err
	}
	return res, nil
}
