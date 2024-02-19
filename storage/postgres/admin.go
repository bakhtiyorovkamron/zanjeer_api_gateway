package postgres

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
)

func (p *postgresRepo) CreateAdmin(req models.Admin) error {

	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	if req.Login == "" || req.Password == "" {
		return fmt.Errorf("login and password are required")
	}

	result, err := p.Db.Db.Exec("insert into admins (id,login,password,type) values ($1,$2,$3,$4)", uuid.String(), req.Login, req.Password, "admin")
	if err != nil {
		return err
	}
	if ok, err := result.RowsAffected(); err != nil || ok == 0 {
		return fmt.Errorf("error while creating admin")
	}

	return nil
}
