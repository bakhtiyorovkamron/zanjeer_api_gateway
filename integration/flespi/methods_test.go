package flespi

import (
	"fmt"
	"testing"
)

func TestGetTelementary(t *testing.T) {
	data, err := GetTelementary()
	if err != nil {
		panic(err)
	}
	for _, device := range data.Result {
		fmt.Println(device.Telemetry.Position)
	}
}
