package backend

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/benspotatoes/extrawatch/models"
	_ "github.com/lib/pq"
)

type Config struct {
	ConnectOpts string
}

type Backend interface {
	IndexMatch(ctx context.Context, params *models.IndexParams) ([]*models.Match, error)
	GetMatch(ctx context.Context, matchID string) (*models.Match, error)
	InsertMatch(ctx context.Context, params *models.InsertParams) error
	UpdateMatch(ctx context.Context, params *models.UpdateParams) error
	DeleteMatch(ctx context.Context, matchID string) error
}

type backendImpl struct {
	db   *sql.DB
	psql squirrel.StatementBuilderType
}

const (
	driver = "postgres"
)

func NewBackend(conf *Config) (Backend, error) {
	db, err := sql.Open(driver, conf.ConnectOpts)
	if err != nil {
		return nil, err
	}
	return &backendImpl{db: db, psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}, nil
}
