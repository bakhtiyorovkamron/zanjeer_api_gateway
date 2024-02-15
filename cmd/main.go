package main

import (
	"github.com/Projects/zanjeer_api_gateway/api"
)

func main() {
	r := api.New()
	r.Run(":8080")
}
