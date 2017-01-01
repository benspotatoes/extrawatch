package backend

import (
	"context"
)

func (b *backendImpl) DeletePlayer(ctx context.Context, playerID string) error {
	_, err := b.deletePlayerQuery(b.parseID(playerID)).RunWith(b.db).Exec()
	return err
}
