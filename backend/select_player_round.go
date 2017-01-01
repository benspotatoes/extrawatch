package backend

import (
	"context"
	"errors"
	"log"

	"github.com/benspotatoes/extrawatch/models"
)

var (
	errMismatchedIDs = errors.New("mismatched player or round id in lookup")
)

func (b *backendImpl) SelectPlayerRound(ctx context.Context, playerID, roundID string) (*models.PlayerRound, error) {
	pid := b.parseID(playerID)
	rid := b.parseID(roundID)
	rows, err := b.selectPlayerRoundQuery(b.parseID(playerID), b.parseID(roundID), 0, 0, "").RunWith(b.db).Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		if deferErr := rows.Close(); deferErr != nil {
			log.Printf("error closing rows: %q", deferErr)
		}
	}()

	var heroes []string
	for rows.Next() {
		var rawPlayerID string
		var rawRoundID string
		var enum int
		err = rows.Scan(&rawPlayerID, &rawRoundID, &enum)
		if err != nil {
			return nil, err
		}
		if rawPlayerID != pid || rawRoundID != rid {
			return nil, errMismatchedIDs
		}
		heroes = append(heroes, models.EnumToHero(enum))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	playerRound := &models.PlayerRound{
		PlayerID: playerID,
		RoundID:  roundID,
		Heroes:   heroes,
	}
	return playerRound, nil
}
