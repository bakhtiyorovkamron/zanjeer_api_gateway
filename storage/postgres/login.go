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
		resp   models.LoginResponse
		status bool
	)
	fmt.Println("req", req)
	data, err := p.Db.Db.Query("select login,password,id,type,created_at,status from admins where login = $1", req.Login)
	if err != nil {
		return resp, err
	}
	for data.Next() {
		err = data.Scan(&res, &password, &resp.Admin.Id, &resp.Admin.Type, &resp.Admin.CreatedAt, &status)
		if err != nil {
			return resp, err
		}
	}
	if res == "" {
		return resp, errors.New("Invalid login")
	}
	if !status {
		return resp, errors.New("Forbidden")
	}

	err = validator.VerifyPassword(req.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("err :", err)
		return resp, errors.New("Invalid login or password mismatch")
	}
	token, err := validator.GenerateToken(resp.Admin.Id, resp.Admin.Type)
	if err != nil {
		return resp, err
	}
	resp.Admin.Login = res
	resp.Token = token
	return resp, nil

}
