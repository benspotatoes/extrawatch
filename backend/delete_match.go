package backend

import (
	"context"
)

const ()

func (b *backendImpl) DeleteMatch(ctx context.Context, matchID string) error {
	_, err := b.deleteMatchQuery(matchID).RunWith(b.db).Exec()
	return err
}
