package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (p *postgresRepo) CreateAdmin(req models.Admin) (models.Admin, error) {

	var admin models.Admin

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return admin, err
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return admin, err
	}
	if req.Login == "" || req.Password == "" {
		return admin, fmt.Errorf("login and password are required")
	}

	result := p.Db.Db.QueryRow("insert into admins (id,login,password,type,first_name,last_name,phone) values ($1,$2,$3,$4,$5,$6,$7) returning id,login,created_at", uuid.String(), req.Login, hashedPassword, "admin", req.Firstname, req.Lastname, req.Phone)
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
	rows, err := p.Db.Db.Query("select id,login,created_at,type from admins where (id = $1 or $1 = '') limit $2 offset $3", req.Id, req.Limit, req.Page)
	if err != nil {
		return admins, err
	}
	for rows.Next() {
		err = rows.Scan(&admin.Id, &admin.Login, &admin.CreatedAt, &admin.Type)
		if err != nil {
			return admins, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}
