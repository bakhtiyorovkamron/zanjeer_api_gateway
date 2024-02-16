package postgres

type PostgresI interface {
	Login(login, password string) (string, error)
}
