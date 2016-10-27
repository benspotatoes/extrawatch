package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) SelectMatch(ctx context.Context, matchID string) (*models.Match, error) {
	var rawID string
	var mapEnum int
	match := &models.Match{}
	rows := b.selectMatchQuery(matchID, 0, 0, "").RunWith(b.db).QueryRow()
	err := rows.Scan(&rawID, &mapEnum, &match.Win, &match.RankDiff, &match.EndingRank, &match.Placement, &match.PlayedOn)
	if err != nil {
		return nil, err
	}

	match.ID = b.parseID(rawID)
	match.Map = models.EnumToMap(mapEnum)

	return match, nil
}
