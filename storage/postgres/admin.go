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
