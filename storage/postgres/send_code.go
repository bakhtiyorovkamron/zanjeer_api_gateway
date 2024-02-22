package postgres

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/google/uuid"
)

func (p *postgresRepo) CreateOTP(req models.SmsOtp) (models.Sms, error) {
	var (
		res models.Sms
	)
	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}
	err = p.Db.Db.QueryRow("insert into sms (id,phone,code) values ($1,$2,$3) returning phone,id", uuid, req.Phone, req.Code).Scan(&res.Phone, &res.SmsId)
	if err != nil {
		fmt.Println("Error while inserting", err)
		return res, err
	}
	return res, nil
}
func (p *postgresRepo) ConfirmOTP(req models.ConfirmOTP) error {
	return p.Db.Db.QueryRow("update sms set confirm = true,code=$2 where id=$1 returning phone", req.SmsId, req.Code).Scan(req.Phone)
}
