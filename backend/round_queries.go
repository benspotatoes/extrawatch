package backend

import (
	"github.com/Masterminds/squirrel"
	"github.com/benspotatoes/extrawatch/models"
)

var (
	roundCols = []string{
		roundIDCol,
		roundMatchIDCol,
		roundCountCol,
		roundModeCol,
		roundTimeLeftCol,
		roundPercentDiffCol,
		roundPointsTakenCol,
		roundNotesCol,
	}
)

const (
	roundTable          = "rounds"
	roundIDCol          = "id"
	roundMatchIDCol     = "match_id"
	roundCountCol       = "count"
	roundModeCol        = "mode"
	roundTimeLeftCol    = "time_left"
	roundPercentDiffCol = "percent_diff"
	roundPointsTakenCol = "points_taken"
	roundNotesCol       = "notes"
)

func (b *backendImpl) selectRoundQuery(roundID string, limit, offset int, filter string) squirrel.SelectBuilder {
	base := b.psql.Select(roundCols...).From(roundTable)
	if roundID != "" {
		base = base.Where("id = ?", roundID)
	}
	if limit != 0 {
		base = base.Limit(uint64(limit))
	}
	if offset != 0 {
		base = base.Offset(uint64(offset))
	}
	if filter != "" {
		// base = base.Where(pred, args)
	}
	return base
}

func (b *backendImpl) insertRoundQuery(params *models.Round) (string, squirrel.InsertBuilder) {
	id := b.newID()
	mID := b.parseID(params.MatchID)
	return b.buildID(roundIDPrefix, id), b.psql.Insert(roundTable).
		Columns(roundCols...).
		Values(id, mID, params.Count, models.ModeToEnum(params.Mode), params.Result.TimeLeft, params.Result.PercentDiff, params.Result.PointsTaken, params.Notes)
}

// TODO - Should we check for blank values before updating a row?
func (b *backendImpl) updateRoundQuery(roundID string, params *models.Round) squirrel.UpdateBuilder {
	return b.psql.Update(roundTable).
		Set(roundCountCol, params.Count).
		Set(roundModeCol, models.ModeToEnum(params.Mode)).
		Set(roundTimeLeftCol, params.Result.TimeLeft).
		Set(roundPercentDiffCol, params.Result.PercentDiff).
		Set(roundPointsTakenCol, params.Result.PointsTaken).
		Set(roundNotesCol, params.Notes).
		// Set(roundPlayedOnCol, params.PlayedOn).
		Where("id = ?", roundID)
}

func (b *backendImpl) deleteRoundQuery(roundID string) squirrel.DeleteBuilder {
	return b.psql.Delete(roundTable).Where("id = ?", roundID)
}
