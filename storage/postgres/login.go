package postgres

import "fmt"

func (p *postgresRepo) Login(login, password string) (string, error) {

	var res string

	data, err := p.Db.Db.Query("select 'hi , kamron '")
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

	return res, err
}
