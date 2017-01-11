package backend

import (
	"context"
	"log"

	"github.com/benspotatoes/extrawatch/models"
)

const (
	maxPlayerLimit = 100
)

func (b *backendImpl) IndexPlayer(ctx context.Context, filter string) ([]*models.Player, error) {
	var players []*models.Player
	rows, err := b.selectPlayerQuery("", maxPlayerLimit, 0, filter).RunWith(b.db).Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close sql connection %q", err.Error())
		}
	}()

	for rows.Next() {
		var rawID string
		player := &models.Player{}

		if err = rows.Scan(&rawID, &player.Name); err != nil {
			return players, err
		}

		player.ID = b.buildID(playerIDPrefix, rawID)
		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		return players, err
	}

	return players, nil
}
