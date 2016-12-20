package backend

import (
	"github.com/Masterminds/squirrel"
	"github.com/benspotatoes/extrawatch/models"
)

var (
	heroCols = []string{
		playerRoundsPlayerIDCol,
		playerRoundsRoundIDCol,
		playerRoundsHeroEnumCol,
	}
)

const (
	playerRoundsTable       = "player_rounds"
	playerRoundsPlayerIDCol = "player_id"
	playerRoundsRoundIDCol  = "round_id"
	playerRoundsHeroEnumCol = "hero"

	heroesDelim = ","
)

func (b *backendImpl) selectPlayerRoundQuery(playerID, roundID string, limit, offset int, filter string) squirrel.SelectBuilder {
	base := b.psql.Select(heroCols...).From(playerRoundsTable)
	if playerID != "" {
		base = base.Where("player_id = ? AND round_id = ?", playerID, roundID)
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

func (b *backendImpl) insertPlayerRoundQuery(playerID, roundID string, hero string) squirrel.InsertBuilder {
	return b.psql.Insert(playerRoundsTable).
		Columns(heroCols...).
		Values(playerID, roundID, models.HeroToEnum(hero))
}

// func (b *backendImpl) updatePlayerRoundQuery(playerID, roundID, hero string) squirrel.UpdateBuilder {
// 	return b.psql.Update(playerRoundsTable).
// 		Set(heroEnumCol, models.HeroToEnum(hero)).
// 		Where("player_id = ? AND round_id = ?", playerID, roundID)
// }

func (b *backendImpl) deletePlayerRoundQuery(playerID, roundID string) squirrel.DeleteBuilder {
	return b.psql.Delete(playerRoundsTable).Where("player_id = ? AND round_id = ?", playerID, roundID)
}
