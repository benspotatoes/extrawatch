package backend

import (
	"context"
	"log"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) IndexMatch(ctx context.Context, limit, offset int, filter string) ([]*models.Match, error) {
	var matches []*models.Match
	rows, err := b.selectMatchQuery("", limit, offset, filter).RunWith(b.db).Query()
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
		var mapEnum int
		match := &models.Match{}

		if err = rows.Scan(&rawID, &mapEnum, &match.Win, &match.RankDiff, &match.EndingRank, &match.Placement, &match.PlayedOn); err != nil {
			return matches, err
		}

		match.ID = b.buildID(matchIDPrefix, rawID)
		match.Map = models.EnumToMap(mapEnum)
		matches = append(matches, match)
	}
	if err = rows.Err(); err != nil {
		return matches, err
	}

	return matches, nil
}
