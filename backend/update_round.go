package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) UpdateRound(ctx context.Context, matchID string, params *models.Round) error {
	_, err := b.updateRoundQuery(matchID, params).RunWith(b.db).Exec()
	return err
}
