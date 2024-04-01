package integration

import (
	"github.com/Projects/zanjeer_api_gateway/integration/flespi"
)

type Integration interface {
	FlespiIntegr() flespi.Flespi
}
type IntegrationP struct {
	Flespi flespi.Flespi
}

func (f *IntegrationP) FlespiIntegr() flespi.Flespi {
	return f.Flespi
}
func New() Integration {
	return &IntegrationP{}
}
