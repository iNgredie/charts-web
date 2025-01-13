package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/iNgredie/charts-web/config"
	"github.com/iNgredie/charts-web/pkg/http_server"
)

type Dependencies struct {
	// Adapters
	//Postgres *postgres.Pool

	// Controllers
	RouterHTTP *chi.Mux
}

func Run(
	ctx context.Context,
	c config.Config,
) (err error) {
	var deps Dependencies

	// Controllers

	//deps.RouterHTTP = router.New()

	httpServer := http_server.New(deps.RouterHTTP, c.HTTP.Port)
	defer httpServer.Close()

	return nil
}
