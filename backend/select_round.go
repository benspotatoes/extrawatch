package backend

import (
	"context"

	"github.com/benspotatoes/extrawatch/models"
)

func (b *backendImpl) SelectRound(ctx context.Context, roundID string) (*models.Round, error) {
	var rawID string
	var modeEnum int
	round := &models.Round{}
	result := &models.Result{}
	rows := b.selectRoundQuery(roundID, 0, 0, "").RunWith(b.db).QueryRow()
	err := rows.Scan(&rawID, &round.Count, &modeEnum, &result.TimeLeft, &result.PercentDiff, &result.PointsTaken, &round.Notes)
	if err != nil {
		return nil, err
	}

	round.ID = b.parseID(rawID)
	round.Mode = models.EnumToMode(modeEnum)
	round.Result = result

	return round, nil
}
