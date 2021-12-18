package main_test

import (
	"net/http"
	"testing"

	"github.com/acme-corp-tech/brick"
	"github.com/acme-corp-tech/brick/test"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/nethttp"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/service"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/storage"
	"github.com/bool64/dbdog"
	"github.com/stretchr/testify/require"
)

func TestFeatures(t *testing.T) {
	var cfg service.Config

	test.RunFeatures(t, "", &cfg, func(tc *test.Context) (*brick.BaseLocator, http.Handler) {
		cfg.ServiceName = service.Name

		sl, err := infra.NewServiceLocator(cfg)
		require.NoError(t, err)

		tc.Database.Instances[dbdog.DefaultDatabase] = dbdog.Instance{
			Tables: map[string]interface{}{
				storage.GreetingsTable: new(storage.GreetingRow),
			},
		}

		return sl.BaseLocator, nethttp.NewRouter(sl)
	})

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
