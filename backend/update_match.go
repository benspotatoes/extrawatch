package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) UpdateMatch(ctx context.Context, matchID string, params *models.Match) error {
	_, err := b.updateMatchQuery(matchID, params).RunWith(b.db).Exec()
	return err
}
