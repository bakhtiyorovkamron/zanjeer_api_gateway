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

func (p *postgresRepo) EditStatus(req models.EditAdminsResponse) error {
	return p.Db.Db.QueryRow("update admins set status=$2  where id=$1 returning status", req.Id, req.Status).Scan(&req.Status)
}

func (p *postgresRepo) GetAdmins(req models.GetAdminsRequest) (models.GetAdminsResponse, error) {
	var (
		admins models.GetAdminsResponse
	)

	rows, err := p.Db.Db.Query(`select id,login,created_at,type,coalesce(first_name,''),coalesce(last_name,''),coalesce(phone,''),(
		select count(*) from admins where  (first_name ilike '%' || $1 || '%' or $1='')
		AND ($4='' OR id = $4)
	) as count,
	coalesce(status,false)
	from admins 
	where  (first_name ilike '%' || $1 || '%' or $1='')
	AND (id=$4 OR $4='')
	order by created_at desc
	limit $2 offset $3`, req.Firstname, req.Limit, req.Limit*(req.Page-1), req.Id)
	if err != nil {
		return admins, err
	}
	for rows.Next() {
		var admin models.Admin

		err = rows.Scan(&admin.Id, &admin.Login, &admin.CreatedAt, &admin.Type, &admin.Firstname, &admin.Lastname, &admin.Phone, &admins.Count,&admin.Status)
		if err != nil {
			return admins, err
		}
		admins.Admins = append(admins.Admins, admin)
	}
	return admins, nil
}
