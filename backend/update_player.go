package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) UpdatePlayer(ctx context.Context, playerID string, params *models.Player) error {
	_, err := b.updatePlayerQuery(b.parseID(playerID), params).RunWith(b.db).Exec()
	return err
}
