package backend

import (
	"github.com/Masterminds/squirrel"
	"github.com/benspotatoes/extrawatch/models"
)

var (
	playerCols = []string{
		playerIDCol,
		playerNameCol,
	}
)

const (
	playerTable   = "players"
	playerIDCol   = "id"
	playerNameCol = "name"
)

func (b *backendImpl) selectPlayerQuery(playerID string, limit, offset int, filter string) squirrel.SelectBuilder {
	base := b.psql.Select(playerCols...).From(playerTable)
	if playerID != "" {
		base = base.Where("id = ?", playerID)
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

func (b *backendImpl) insertPlayerQuery(params *models.Player) (string, squirrel.InsertBuilder) {
	id := b.newID()
	return b.buildID(playerIDPrefix, id), b.psql.Insert(playerTable).
		Columns(playerCols...).
		Values(id, params.Name)
}

func (b *backendImpl) updatePlayerQuery(playerID string, params *models.Player) squirrel.UpdateBuilder {
	return b.psql.Update(playerTable).
		Set(playerNameCol, params.Name).
		Where("id = ?", playerID)
}

func (b *backendImpl) deletePlayerQuery(playerID string) squirrel.DeleteBuilder {
	return b.psql.Delete(playerTable).Where("id = ?", playerID)
}
