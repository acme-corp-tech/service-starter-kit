package service

import (
	"github.com/acme-corp-tech/service-starter-kit/internal/domain/greeting"
)

// GreetingMakerProvider is a service provider.
type GreetingMakerProvider interface {
	GreetingMaker() greeting.Maker
}
