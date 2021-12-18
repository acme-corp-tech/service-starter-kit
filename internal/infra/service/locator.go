package service

import (
	"github.com/acme-corp-tech/brick"
)

// Locator defines application resources.
type Locator struct {
	*brick.BaseLocator

	GreetingMakerProvider
}
