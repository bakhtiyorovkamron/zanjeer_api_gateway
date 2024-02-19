package postgres

import "github.com/Projects/zanjeer_api_gateway/models"

type PostgresI interface {
	Login(req models.Login) (string, error)
}
