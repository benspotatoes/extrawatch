package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) GetMatch(ctx context.Context, matchID string) (*models.Match, error) {
	return nil, errNotImplemented
}
