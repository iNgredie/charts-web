package main

import (
	"context"
	"github.com/iNgredie/charts-web/config"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}
}
