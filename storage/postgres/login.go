package postgres

import (
	"errors"
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/validator"
)

func (p *postgresRepo) Login(req models.Login) (string, error) {

	var res string

	data, err := p.Db.Db.Query("select login from admins where login = $1 and password = $2 and type=$3", req.Login, req.Password, req.Type)
	if err != nil {
		fmt.Println("Error while querying", err)
		return "", err
	}
	for data.Next() {
		err = data.Scan(&res)
		if err != nil {
			fmt.Println("Error while scanning", err)
			return "", err
		}
	}
	if res == "" {
		return "", errors.New("Invalid login")
	}
	token, err := validator.GenerateToken(req.Login, req.Type)
	if err != nil {
		return "", err
	}
	return token, nil

}
