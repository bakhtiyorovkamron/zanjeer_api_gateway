package main

import (
	"github.com/Projects/zanjeer_api_gateway/api"
	"github.com/Projects/zanjeer_api_gateway/config"
)

func main() {
	cfg := config.Load()

	r := api.New(cfg)
	r.Run(":7777")
}
