package postgres

import (
	"fmt"
	"testing"

	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/pkg/db"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

var cfg = config.Load()

type DriverTest struct {
	Phone     string `fake:"{phone}"`
	Firstname string `fake:"{firstname}"`
	Lastname  string `fake:"{lastname}"`
}

func TestCreateDrivers(t *testing.T) {
	db, err := db.New(cfg)
	if err != nil {
		fmt.Println("Failed to create")
	} else {
		fmt.Println("err :", err)
	}

	var driver DriverTest

	for i := 0; i < 1000; i++ {

		uuid, _ := uuid.NewUUID()

		err = gofakeit.Struct(&driver)
		if err != nil {
			panic(err)
		}

		_, err = db.Db.Query("insert into drivers (id,phone,first_name,last_name) values ($1,$2,$3,$4) returning phone,id", uuid, driver.Phone, driver.Firstname, driver.Lastname)
		if err != nil {
			panic(err)
		}
	}

}
