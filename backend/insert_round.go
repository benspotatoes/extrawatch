package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) InsertRound(ctx context.Context, params *models.Round) (string, error) {
	id, query := b.insertRoundQuery(params)
	_, err := query.RunWith(b.db).Exec()
	return id, err
}
