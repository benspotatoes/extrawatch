package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) UpdateRound(ctx context.Context, roundID string, params *models.Round) error {
	_, err := b.updateRoundQuery(b.parseID(roundID), params).RunWith(b.db).Exec()
	return err
}
