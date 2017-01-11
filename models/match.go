package models

import "time"

var (
	validMaps map[string]bool
	maps      = []string{
		// Escort
		"watchpoint",
		"dorado",
		"route66",
		// Assault
		"hanamura",
		"anubis",
		"volskaya",
		// Hybrid
		"kings",
		"numbani",
		"hollywood",
		"eichenwalde",
		// Control
		"lijiang",
		"nepal",
		"ilios",
		"oasis"
	}
	types = []string{
		"Escort",
		"Assault",
		"Hybrid",
		"Control",
	}
)

const (
	minRankDiff = 500
	minRank     = 0
	maxRank     = 5000
)

func init() {
	validMaps = make(map[string]bool, len(maps))
	for _, name := range maps {
		validMaps[name] = true
	}
}

type Match struct {
	ID  string `json:"id"`
	Map string `json:"map,omitempty"`
	// Rounds     []*Round `json:"rounds"`
	Win        int        `json:"win,omitempty"`
	RankDiff   int        `json:"rank_diff,omitempty"`
	EndingRank int        `json:"ending_rank,omitempty"`
	Placement  bool       `json:"placement,omitempty"`
	PlayedOn   *time.Time `json:"played_on,omitempty"`
}

func (m *Match) Validate() error {
	// Map
	if !validMaps[m.Map] {
		return errInvalidMap
	}

	// // Rounds
	// if len(m.Rounds) < 2 {
	// 	return errInvalidMatchRounds
	// }
	// for _, round := range m.Rounds {
	// 	// Round
	// 	if roundErr := round.validate(); roundErr != nil {
	// 		return roundErr
	// 	}
	// }

	// Win
	if m.Win < -1 || m.Win > 1 {
		return errInvalidMatchResult
	}

	// RankDiff
	if m.RankDiff > minRankDiff || m.RankDiff < -minRankDiff {
		return errInvalidRankDiff
	}

	// EndingRank
	if m.EndingRank < minRank || m.EndingRank > maxRank {
		return errInvalidEndingRank
	}

	return nil
}

func MapToEnum(name string) int {
	for i, m := range maps {
		if m == name {
			return i
		}
	}
	return -1
}

func EnumToMap(enum int) string {
	return maps[enum]
}
