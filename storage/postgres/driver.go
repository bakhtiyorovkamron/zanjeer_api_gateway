package postgres

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
)

func (p *postgresRepo) CreateDriver(req models.Driver) (models.Driver, error) {
	var (
		res models.Driver
	)
	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}
	err = p.Db.Db.QueryRow("insert into drivers (id,phone,first_name,last_name) values ($1,$2,$3,$4) returning phone,id,first_name,last_name", uuid, req.Phone, req.Firstname, req.Lastname).Scan(&res.Phone, &res.Id, &res.Firstname, &res.Lastname)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (p *postgresRepo) GetDriverInfo(id string) (models.Driver, error) {
	var (
		res models.Driver
	)
	data, err := p.Db.Db.Query("SELECT id,phone,first_name,last_name FROM drivers WHERE id = $1", id)
	if err != nil {
		return res, err
	}
	fmt.Println("data.Err() :", data.Err())
	for data.Next() {
		if err != data.Scan(&res.Id, &res.Phone, &res.Firstname, &res.Lastname) {
			return res, err
		}
	}
	return res, nil
}
func (p *postgresRepo) UpdateDriverInfo(req models.Driver) (models.Driver, error) {
	var res models.Driver

	err := p.Db.Db.QueryRow(`
	WITH d AS (
		SELECT * FROM "drivers" WHERE id=$1
	)
	UPDATE "drivers" SET 
	phone = (
		CASE 
			WHEN length($2) > 0 THEN $2
			ELSE d.phone
		END
	),
	first_name = (
		CASE 
			WHEN length($3) > 0 THEN $3
			ELSE d.first_name
		END
	),
	last_name = (
		CASE
			WHEN length($4) > 0 THEN $4
			ELSE d.last_name
		END
	)
	FROM d
	WHERE drivers.id = d.id
	RETURNING drivers.id,drivers.phone,drivers.first_name,drivers.last_name
	`, req.Id, req.Phone, req.Firstname, req.Lastname).Scan(&res.Id, &res.Phone, &res.Firstname, &res.Lastname)
	if err != nil {
		fmt.Println("err :", err)
		if strings.ContainsAny(err.Error(), "no rows found") {
			return res, fmt.Errorf("driver does not exist")
		}
		return res, err
	}
	return res, err

}
func (p *postgresRepo) DeleteDriver(id string) error {
	query := p.Db.Builder.Delete("drivers").Where(squirrel.Eq{"id": id})
	_, err := query.RunWith(p.Db.Db).Exec()
	return err
}
func (p *postgresRepo) GetDriverList(limit, offset int64) (models.DriverList, error) {

	var (
		drivers models.DriverList
		count   int32
	)

	query := `SELECT id,
					phone,
					coalesce(first_name,''),
					coalesce(last_name,''),
					created_at,
					(select count(*) from drivers) as count
			FROM drivers
			ORDER BY created_at
			LIMIT $1
			OFFSET $2
			`
	data, err := p.Db.Db.Query(query, limit, limit*(offset-1))
	if err != nil {
		return drivers, err
	}
	for data.Next() {
		var driver models.Driver
		if err := data.Scan(&driver.Id, &driver.Phone, &driver.Firstname, &driver.Lastname, &driver.CreatedAt, &count); err != nil {
			return drivers, err
		}
		drivers.Drivers = append(drivers.Drivers, driver)
	}
	drivers.Count = count

	return drivers, nil
}
func (p *postgresRepo) SearchDriver(req models.DriverSearchRequest) (models.DriverList, error) {

	var (
		drivers models.DriverList
		count   int
	)

	query := `SELECT id,
					phone,
					first_name,
					last_name,
					created_at,
					(
						SELECT count(*)
			FROM drivers
			WHERE ($1='' OR phone ilike '%' || $1 || '%')
			AND   ($2='' OR first_name ilike '%'  || $2 || '%')
			AND   ($3='' OR last_name  ilike '%'  || $3 || '%')
			) as count
			FROM drivers
			WHERE ($1='' OR phone ilike '%' || $1 || '%')
			AND   ($2='' OR first_name ilike '%'  || $2 || '%')
			AND   ($3='' OR last_name  ilike '%'  || $3 || '%')
			ORDER BY created_at
			LIMIT $4
			OFFSET $5
			`
	data, err := p.Db.Db.Query(query, req.Phone, req.Firstname, req.Lastname, req.Limit, req.Limit*(req.Offset-1))
	if err != nil {
		return drivers, err
	}
	for data.Next() {
		var driver models.Driver
		if err := data.Scan(&driver.Id, &driver.Phone, &driver.Firstname, &driver.Lastname, &driver.CreatedAt, &count); err != nil {
			return drivers, err
		}
		drivers.Drivers = append(drivers.Drivers, driver)
	}
	drivers.Count = int32(count)

	return drivers, nil
}
func (p *postgresRepo) VerifyDriver(phone string) (models.Driver, error) {
	var res models.Driver

	err := p.Db.Db.QueryRow(`
	UPDATE "drivers" SET 
	verified = true WHERE phone = $1
	RETURNING id,phone,first_name,last_name,verified
	`, phone).Scan(&res.Id, &res.Phone, &res.Firstname, &res.Lastname, &res.Verified)
	if err != nil {
		if strings.ContainsAny(err.Error(), "no rows found") {
			return res, fmt.Errorf("driver does not exist")
		}
		return res, err
	}
	return res, err

}
