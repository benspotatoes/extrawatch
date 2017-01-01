package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) InsertPlayer(ctx context.Context, params *models.Player) (string, error) {
	id, query := b.insertPlayerQuery(params)
	_, err := query.RunWith(b.db).Exec()
	return id, err
}
