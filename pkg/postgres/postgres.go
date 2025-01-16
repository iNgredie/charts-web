package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Config struct {
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
	Db       string `envconfig:"POSTGRES_DB" required:"true"`
}

type Pool struct {
	*pgxpool.Pool
}

func New(
	ctx context.Context,
	c Config,
) (*Pool, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Db,
	)
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	return &Pool{Pool: pool}, nil
}

func (p Pool) Close() {
	p.Pool.Close()

	log.Info().Msg("Postgres closed")
}
