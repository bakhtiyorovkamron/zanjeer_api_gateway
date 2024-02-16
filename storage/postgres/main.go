package postgres

import (
	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/pkg/db"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
)

type postgresRepo struct {
	Db  *db.Postgres
	Log *logger.Logger
	Cfg config.Config
}

func New(db *db.Postgres, log *logger.Logger, cfg config.Config) PostgresI {
	return &postgresRepo{
		Db:  db,
		Log: log,
		Cfg: cfg,
	}
}
