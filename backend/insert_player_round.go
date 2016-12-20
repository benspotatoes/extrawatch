package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) InsertPlayerRound(ctx context.Context, params *models.PlayerRound) error {
	for _, hero := range params.Heroes {
		query := b.insertPlayerRoundQuery(b.parseID(params.PlayerID), b.parseID(params.RoundID), hero)
		_, err := query.RunWith(b.db).Exec()
		// TODO - Make this a transaction
		if err != nil {
			return err
		}
	}
	return nil
}
