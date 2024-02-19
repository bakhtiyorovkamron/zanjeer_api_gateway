package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
)

func (p *postgresRepo) CreateAdmin(req models.Admin) (models.Admin, error) {

	var admin models.Admin

	uuid, err := uuid.NewUUID()
	if err != nil {
		return admin, err
	}
	if req.Login == "" || req.Password == "" {
		return admin, fmt.Errorf("login and password are required")
	}

	result := p.Db.Db.QueryRow("insert into admins (id,login,password,type) values ($1,$2,$3,$4) returning id,login,created_at", uuid.String(), req.Login, req.Password, "admin")
	if err != nil {
		return admin, err
	}
	var createdAt sql.NullString
	err = result.Scan(&admin.Id, &admin.Login, &createdAt)
	if err != nil {
		return admin, err
	}
	admin.CreatedAt = createdAt.String

	return admin, nil
}
func (p *postgresRepo) GetAdmins(req models.GetAdmins) ([]models.Admin, error) {
	var admins []models.Admin
	var admin models.Admin
	fmt.Println("limit", req.Limit, "page", req.Page)
	rows, err := p.Db.Db.Query("select id,login,created_at from admins limit $1 offset $2", req.Limit, req.Page)
	if err != nil {
		return admins, err
	}
	for rows.Next() {
		err = rows.Scan(&admin.Id, &admin.Login, &admin.CreatedAt)
		if err != nil {
			return admins, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}
