package postgres

import (
	"errors"
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

func (p *postgresRepo) Login(req models.Login) (models.LoginResponse, error) {

	var (
		res,
		password string
		resp models.LoginResponse
	)

	data, err := p.Db.Db.Query("select login,password,id,type,created_at from admins where login = $1", req.Login)
	if err != nil {
		fmt.Println("Error while querying", err)
		return resp, err
	}
	for data.Next() {
		err = data.Scan(&res, &password, &resp.Admin.Id, &resp.Admin.Type, &resp.Admin.CreatedAt)
		if err != nil {
			fmt.Println("Error while scanning", err)
			return resp, err
		}
	}
	if res == "" {
		return resp, errors.New("Invalid login")
	}

	err = validator.VerifyPassword(req.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return resp, errors.New("Invalid login or password mismatch")
	}
	fmt.Println("Admin type", resp.Admin.Type)
	token, err := validator.GenerateToken(resp.Admin.Id, resp.Admin.Type)
	if err != nil {
		return resp, err
	}
	resp.Admin.Login = res
	resp.Token = token
	return resp, nil

}
