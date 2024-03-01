package postgres

import (
	"database/sql"

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
func (p *postgresRepo) GetDeviceTypeList(req models.GetDeviceTypeListRequest) ([]models.DeviceType, error) {
	var (
		res []models.DeviceType
	)
	query := `SELECT id,
					name
			FROM device_type
			WHERE name ilike  '%' || $1 || '%'
			`
	data, err := p.Db.Db.Query(query, req.Name)
	if err != nil {
		return res, err
	}
	for data.Next() {
		var d models.DeviceType
		if err := data.Scan(&d.Id, &d.Name); err != nil {
			return res, err
		}
		res = append(res, d)
	}
	return res, nil
}

func (p *postgresRepo) CreateDevice(req models.CreateDeviceRequest) (models.CreateDeviceResponse, error) {
	var (
		res       models.CreateDeviceResponse
		createdAt sql.NullString
	)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}
	err = p.Db.Db.QueryRow(`insert into devices 
	(
		id,
		name,
		type,
		address,
		imei,
		driver
	) values ($1,$2,$3,$4,$5,$6)
	returning id,name,type,address,imei,driver,created_at`, uuid, req.Name, req.Type, req.IpAddress, req.Imei, req.DriverId).Scan(
		&res.Id,
		&res.Name,
		&res.Type,
		&res.IpAddress,
		&res.Imei,
		&res.DriverId,
		&createdAt,
	)
	if err != nil {
		return res, err
	}
	res.CreatedAt = createdAt.String

	return res, nil
}
