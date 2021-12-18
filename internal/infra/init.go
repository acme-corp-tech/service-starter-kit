package infra

import (
	"context"

	"github.com/acme-corp-tech/brick"
	"github.com/acme-corp-tech/brick/database"
	"github.com/acme-corp-tech/brick/jaeger"
	"github.com/acme-corp-tech/service-starter-kit/internal/domain/greeting"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/schema"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/service"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/storage"
	"github.com/go-sql-driver/mysql"
	"github.com/swaggest/rest/response/gzip"
)

// NewServiceLocator creates application service locator.
func NewServiceLocator(cfg service.Config) (loc *service.Locator, err error) {
	l := &service.Locator{}

	defer func() {
		if err != nil && l != nil && l.LoggerProvider != nil {
			l.CtxdLogger().Error(context.Background(), err.Error())
		}
	}()

	l.BaseLocator, err = brick.NewBaseLocator(cfg.BaseConfig)
	if err != nil {
		return nil, err
	}

	if err = jaeger.Setup(cfg.Jaeger, l.BaseLocator); err != nil {
		return nil, err
	}

	schema.SetupOpenapiCollector(l.OpenAPI)

	l.HTTPServerMiddlewares = append(l.HTTPServerMiddlewares, gzip.Middleware)

	if err = setupStorage(l, cfg.Database); err != nil {
		return nil, err
	}

	l.GreetingMakerProvider = &storage.GreetingSaver{
		Upstream: &greeting.SimpleMaker{},
		Storage:  l.Storage,
	}

	return l, nil
}

func setupStorage(l *service.Locator, cfg database.Config) error {
	c, err := mysql.ParseDSN(cfg.DSN)
	if err != nil {
		return err
	}

	conn, err := mysql.NewConnector(c)
	if err != nil {
		return err
	}

	l.Storage, err = database.SetupStorage(cfg, l.CtxdLogger(), "mysql", conn, storage.Migrations)
	if err != nil {
		return err
	}

	return nil
}
