package postgres

import (
	"fmt"
	"strings"

	"github.com/Projects/zanjeer_api_gateway/models"
)

func (p *postgresRepo) GetDeviceLocation(req models.GetDeviceLocationRequest) ([]models.GetDeviceLocationResponse, error) {

	resp := []models.GetDeviceLocationResponse{}

	data, err := p.Db.Db.Query("SELECT imei,longitude,latitiude,created_at FROM devices_location order by created_at  limit 1")
	if err != nil {
		return resp, err
	}
	for data.Next() {
		var (
			imei                          string
			longitude, latitude           []byte
			longitudeSlice, latitudeSlice = []string{}, []string{}
			createdAt                     string
		)

		if err := data.Scan(
			&imei,
			&longitude,
			&latitude,
			&createdAt,
		); err != nil {
			return resp, err
		}

		longitudeString, latitudeString := string(longitude), string(latitude)

		if longitudeString != "{}" {
			longitudeString = longitudeString[1 : len(longitudeString)-1] // Remove curly braces
			longitudeSlice = strings.Split(longitudeString, ",")
			for i := range longitudeSlice {
				longitudeSlice[i] = strings.TrimSpace(longitudeSlice[i]) // Remove leading/trailing whitespaces
			}
		} else {
			longitudeSlice = []string{} // empty array
		}
		if latitudeString != "" {
			latitudeString = latitudeString[1 : len(latitudeString)-1]
			latitudeSlice = strings.Split(latitudeString, ",")
			for i := range latitudeSlice {
				latitudeSlice[i] = strings.TrimSpace(latitudeSlice[i]) // Remove leading/trailing whitespaces
			}
		} else {
			latitudeSlice = []string{}
		}
		fmt.Println("latitudeSlice: ", latitudeSlice)
		fmt.Println("longitudeSlice: ", longitudeSlice)
		resp = append(resp, models.GetDeviceLocationResponse{
			Imei: imei,
			Time: createdAt,
			Longitude: func() string {
				for i := range longitudeSlice {
					if string(longitudeSlice[i]) != "0" {
						return string(longitudeSlice[i])
					}
				}
				return ""
			}(),
			Latitude: func() string {
				for i := range latitudeSlice {
					if string(latitudeSlice[i]) != "0" {
						return string(latitudeSlice[i])
					}
				}
				return ""
			}(),
		})
		// resp = append(resp, models.GetDeviceLocationResponse{
		// 	Imei:      imei,
		// 	Time:      createdAt,
		// 	Longitude: longitudeSlice,
		// 	Latitude:  latitudeSlice,
		// })
	}

	return resp, nil
}
