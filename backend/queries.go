package backend

import (
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/benspotatoes/extrawatch/models"
	uuid "github.com/satori/go.uuid"
)

var (
	matchCols = []string{
		matchIDCol,
		matchMapCol,
		matchWinCol,
		matchRankDiffCol,
		matchEndingRankCol,
		matchPlacementCol,
		matchPlayedOnCol,
	}
)

const (
	idDelim       = ":"
	matchIDPrefix = "match"

	matchTable         = "matches"
	matchIDCol         = "id"
	matchMapCol        = "map"
	matchWinCol        = "win"
	matchRankDiffCol   = "rank_diff"
	matchEndingRankCol = "ending_rank"
	matchPlacementCol  = "placement"
	matchPlayedOnCol   = "played_on"
)

func (b *backendImpl) newID(prefix string) string {
	return fmt.Sprintf("%s%s%s", prefix, idDelim, uuid.NewV4().String())
}

func (b *backendImpl) buildID(prefix, id string) string {
	return fmt.Sprintf("%s%s%s", prefix, idDelim, id)
}

func (b *backendImpl) parseID(id string) string {
	split := strings.Split(id, idDelim)
	return split[len(split)-1]
}

func (b *backendImpl) selectMatchQuery(matchID string, limit, offset int, filter string) squirrel.SelectBuilder {
	base := b.psql.Select(matchCols...).From(matchTable)
	if matchID != "" {
		base = base.Where("id = ?", b.buildID(matchIDPrefix, matchID))
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

func (b *backendImpl) insertMatchQuery(params *models.Match) (string, squirrel.InsertBuilder) {
	id := b.newID(matchIDPrefix)
	return b.parseID(id), b.psql.Insert(matchTable).
		Columns(matchCols...).
		Values(id, models.MapToEnum(params.Map), params.Win, params.RankDiff, params.EndingRank, params.Placement, time.Now())
}

// TODO - Should we check for blank values before updating a row?
func (b *backendImpl) updateMatchQuery(matchID string, params *models.Match) squirrel.UpdateBuilder {
	return b.psql.Update(matchTable).
		Set(matchMapCol, models.MapToEnum(params.Map)).
		Set(matchWinCol, params.Win).
		Set(matchRankDiffCol, params.RankDiff).
		Set(matchEndingRankCol, params.EndingRank).
		Set(matchPlacementCol, params.Placement).
		// Set(matchPlayedOnCol, params.PlayedOn).
		Where("id = ?", b.buildID(matchIDPrefix, matchID))
}

func (b *backendImpl) deleteMatchQuery(matchID string) squirrel.DeleteBuilder {
	return b.psql.Delete(matchTable).Where("id = ?", b.newID(matchID))
}
