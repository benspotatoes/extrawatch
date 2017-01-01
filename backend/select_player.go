package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) SelectPlayer(ctx context.Context, playerID string) (*models.Player, error) {
	var rawID string
	player := &models.Player{}
	rows := b.selectPlayerQuery(b.parseID(playerID), 0, 0, "").RunWith(b.db).QueryRow()
	err := rows.Scan(&rawID, &player.Name)
	if err != nil {
		return nil, err
	}
	player.ID = playerID
	return player, nil
}
