package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) IndexMatch(ctx context.Context, params *models.IndexParams) ([]*models.Match, error) {
	return []*models.Match{}, errNotImplemented
}
