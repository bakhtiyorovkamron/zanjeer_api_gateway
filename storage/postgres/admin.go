package postgres

import "github.com/Projects/zanjeer_api_gateway/models"

func (p *postgresRepo) CreateAdmin(req models.Admin) error {
	// _, err := p.Db.Db.Exec("insert into admins(login,password,type) values($1,$2,$3)", req.Login, req.Password, req.Type)
	// if err != nil {
	// 	return err
	// }
	return nil
}
