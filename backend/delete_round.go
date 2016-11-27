package backend

import (
	"context"
)

const ()

func (b *backendImpl) DeleteRound(ctx context.Context, roundID string) error {
	_, err := b.deleteRoundQuery(b.parseID(roundID)).RunWith(b.db).Exec()
	return err
}
