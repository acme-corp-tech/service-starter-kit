package service

import (
	"github.com/acme-corp-tech/brick"
	"github.com/acme-corp-tech/brick/database"
	"github.com/acme-corp-tech/brick/jaeger"
)

// Name is the name of this application or service.
const Name = "service-starter-kit"

// Config defines application configuration.
type Config struct {
	brick.BaseConfig

	Database database.Config `split_words:"true"`
	Jaeger   jaeger.Config   `split_words:"true"`
}
