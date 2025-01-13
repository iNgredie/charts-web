package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/iNgredie/charts-web/config"
	"github.com/iNgredie/charts-web/pkg/http_server"
	"github.com/iNgredie/charts-web/pkg/router"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
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

	deps.RouterHTTP = router.New()

	httpServer := http_server.New(deps.RouterHTTP, c.HTTP.Port)
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	log.Info().Msg("App started")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		log.Info().Msg("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		log.Error().Err(err).Msg("App got notify: httpServer.Notify")
	}

	log.Info().Msg("App is stopping...")
}
