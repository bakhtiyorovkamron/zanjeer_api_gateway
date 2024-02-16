package postgres

func (p *postgresRepo) Login(login, password string) (string, error) {

	var res string

	data, err := p.Db.Db.Query("select 'hi , kamron '")
	if err != nil {
		return "", err
	}
	err = data.Scan(&res)

	return res, err
}
