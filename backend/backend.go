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
	IndexMatch(ctx context.Context, limit, offset int, filter string) ([]*models.Match, error)
	SelectMatch(ctx context.Context, matchID string) (*models.Match, error)
	InsertMatch(ctx context.Context, params *models.Match) (string, error)
	UpdateMatch(ctx context.Context, matchID string, params *models.Match) error
	DeleteMatch(ctx context.Context, matchID string) error

	SelectRound(ctx context.Context, roundID string) (*models.Round, error)
	InsertRound(ctx context.Context, params *models.Round) (string, error)
	UpdateRound(ctx context.Context, roundID string, params *models.Round) error
	DeleteRound(ctx context.Context, roundID string) error

	IndexPlayer(ctx context.Context, filter string) ([]*models.Player, error)
	SelectPlayer(ctx context.Context, playerID string) (*models.Player, error)
	InsertPlayer(ctx context.Context, params *models.Player) (string, error)
	UpdatePlayer(ctx context.Context, playerID string, params *models.Player) error
	DeletePlayer(ctx context.Context, playerID string) error

	SelectPlayerRound(ctx context.Context, playerID, roundID string) (*models.PlayerRound, error)
	InsertPlayerRound(ctx context.Context, params *models.PlayerRound) error
	UpdatePlayerRound(ctx context.Context, playerID, roundID string, params *models.PlayerRound) error
	DeletePlayerRound(ctx context.Context, playerID, roundID string) error
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
