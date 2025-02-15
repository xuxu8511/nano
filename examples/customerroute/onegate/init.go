package onegate

import "github.com/lonng/nano/component"

var (
	// Services All services in master server
	Services = &component.Components{}

	bindService = newRegisterService()
)

func init() {
	Services.Register(bindService)
}
