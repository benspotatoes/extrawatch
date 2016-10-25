package backend

import "github.com/Masterminds/squirrel"

const (
	matchTable = "matches"
)

func (b *backendImpl) deleteMatchQuery(matchID string) squirrel.DeleteBuilder {
	return b.psql.Delete(matchTable).Where("id = ?", matchID)
}
