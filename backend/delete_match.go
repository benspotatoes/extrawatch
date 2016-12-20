package backend

import (
	"context"
)

func (b *backendImpl) DeleteMatch(ctx context.Context, matchID string) error {
	_, err := b.deleteMatchQuery(b.parseID(matchID)).RunWith(b.db).Exec()
	return err
}
