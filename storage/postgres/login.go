package postgres

import (
	"errors"
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

func (p *postgresRepo) Login(req models.Login) (string, error) {

	var (
		res,
		password string
	)

	data, err := p.Db.Db.Query("select login,password from admins where login = $1 ", req.Login)
	if err != nil {
		fmt.Println("Error while querying", err)
		return "", err
	}
	for data.Next() {
		err = data.Scan(&res, &password)
		if err != nil {
			fmt.Println("Error while scanning", err)
			return "", err
		}
	}
	if res == "" {
		return "", errors.New("Invalid login")
	}

	fmt.Println("Password :", password)
	fmt.Println("Req Password:", req.Password)

	err = validator.VerifyPassword(req.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("Invalid login or password mismatch")
	}

	token, err := validator.GenerateToken(req.Login, req.Type)
	if err != nil {
		return "", err
	}
	return token, nil

}
