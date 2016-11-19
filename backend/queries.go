package backend

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	idDelim       = ":"
	matchIDPrefix = "match"
	roundIDPrefix = "round"
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
