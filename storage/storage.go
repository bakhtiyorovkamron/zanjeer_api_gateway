package storage

import (
	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/pkg/db"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
	"github.com/Projects/zanjeer_api_gateway/storage/postgres"
)

type StorageI interface {
	Postgres() postgres.PostgresI
}
type StoragePg struct {
	postgres postgres.PostgresI
}

func (s *StoragePg) Postgres() postgres.PostgresI {
	return s.postgres
}
func New(db *db.Postgres, log *logger.Logger, cfg config.Config) StorageI {
	return &StoragePg{
		postgres: postgres.New(db, log, cfg),
	}
}
