package models

import "errors"

var (
	// Match
	errInvalidMap         = errors.New("invalid map")
	errInvalidMatchResult = errors.New("invalid match result")
	errInvalidMatchRounds = errors.New("invalid match rounds")
	errInvalidRankDiff    = errors.New("invalid rank diff")
	errInvalidEndingRank  = errors.New("invalid ending rank")
	// Player
	errInvalidPlayer = errors.New("invalid player")
	errInvalidHeroes = errors.New("invalid hero count")
	errInvalidHero   = errors.New("invalid hero")
	// Round
	errInvalidRoundCount   = errors.New("invalid round count")
	errInvalidRoundPlayers = errors.New("invalid round player count")
	errInvalidMode         = errors.New("invalid mode")
	errInvalidRoundResult  = errors.New("invalid round result")
	errInvalidNotes        = errors.New("invalid notes")
)
