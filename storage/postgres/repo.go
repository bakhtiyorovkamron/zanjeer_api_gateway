package postgres

import "github.com/Projects/zanjeer_api_gateway/models"

type PostgresI interface {
	Login(req models.Login) (models.LoginResponse, error)
	CreateAdmin(req models.Admin) (models.Admin, error)
	GetAdmins(req models.GetAdmins) ([]models.Admin, error)
	CreateOTP(req models.SmsOtp) (models.Sms, error)
	ConfirmOTP(req models.ConfirmOTP) error

	//Driver represents
	CreateDriver(req models.Driver) (models.Driver, error)
	UpdateDriverInfo(req models.Driver) (models.Driver, error)
	GetDriverInfo(id string) (models.Driver, error)
	DeleteDriver(id string) error
	GetDriverList(limit, offset int64) (models.DriverList, error)

	//Devices
	CreateDeviceType(req models.DeviceType) (models.DeviceType, error)
}
