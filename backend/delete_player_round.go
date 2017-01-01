package backend

import (
	"context"
)

func (b *backendImpl) DeletePlayerRound(ctx context.Context, playerID, roundID string) error {
	_, err := b.deletePlayerRoundQuery(b.parseID(playerID), b.parseID(roundID)).RunWith(b.db).Exec()
	return err
}
