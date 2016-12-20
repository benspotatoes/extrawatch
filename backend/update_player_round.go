package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) UpdatePlayerRound(ctx context.Context, playerID, roundID string, params *models.PlayerRound) error {
	return errNotImplemented
	// _, err := b.updatePlayerRoundQuery(b.parseID(playerID), b.parse(roundID), params).RunWith(b.db).Exec()
	// return err
}
