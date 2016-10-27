package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) InsertMatch(ctx context.Context, params *models.Match) (string, error) {
	id, query := b.insertMatchQuery(params)
	_, err := query.RunWith(b.db).Exec()
	return id, err
}
