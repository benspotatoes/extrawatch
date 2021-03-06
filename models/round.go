package models

var (
	validModes map[string]bool
	modes      = []string{
		// Escort, Assault, Hybrid
		"attack",
		"defend",
		// LijiangTower
		"market",
		"garden",
		"center",
		// Nepal
		"village",
		"shrine",
		"sanctum",
		// Ilios
		"lighthouse",
		"well",
		"ruins",
		// Oasis
		"center",
		"gardens",
		"university",
	}
)

const (
	playerCount = 6
)

func init() {
	validModes = make(map[string]bool, len(modes))
	for _, name := range modes {
		validModes[name] = true
	}
}

type Round struct {
	ID      string `json:"id"`
	MatchID string `json:"match_id,omitempty"`
	Count   int    `json:"count,omitempty"`
	// Players []*Player `json:"players"`
	Mode   string  `json:"mode,omitempty"`
	Result *Result `json:"result,omitempty"`
	Notes  string  `json:"notes,omitempty"`
}

func (r *Round) Validate() error {
	// Match ID
	if r.MatchID == "" {
		return errInvalidMatchID
	}

	// Count
	if r.Count < 0 {
		return errInvalidRoundCount
	}

	// // Players
	// if len(r.Players) != playerCount {
	// 	return errInvalidRoundPlayers
	// }
	// for _, player := range r.Players {
	// 	// Player
	// 	if playErr := player.Validate(); playErr != nil {
	// 		return playErr
	// 	}
	// }

	// Mode
	// TODO - Validate mode with map context
	if !validModes[r.Mode] {
		return errInvalidMode
	}

	// Result
	if resErr := r.Result.Validate(); resErr != nil {
		return resErr
	}

	// Notes
	if len(r.Notes) > 500 {
		return errInvalidNotes
	}

	return nil
}

type Result struct {
	Win         int `json:"win"`
	TimeLeft    int `json:"time_left"`
	PercentDiff int `json:"percent_diff"`
	PointsTaken int `json:"points_taken"`
}

func (r *Result) Validate() error {
	if r.Win < -1 || r.Win > 1 {
		return errInvalidRoundResult
	}

	switch {
	case r.TimeLeft != 0 && r.PercentDiff == 0 && r.PointsTaken == 0:
		// Payload/Hybrid
		// Time remaining means the payload was pushed to the last point
		return nil
	case r.PercentDiff != 0 && r.TimeLeft == 0 && r.PointsTaken == 0:
		// King of the hill
		// Positive percent diff means a win, negative means a loss
		return nil
	case r.PointsTaken >= 0 && r.TimeLeft == 0 && r.PercentDiff == 0:
		// Assault/Defend
		// No points taken with no time left means either the attack was a super
		// disaster or the defense was a great success
		return nil
	default:
		return errInvalidRoundResult
	}
}

func ModeToEnum(name string) int {
	for i, m := range modes {
		if m == name {
			return i
		}
	}
	return -1
}

func EnumToMode(enum int) string {
	return modes[enum]
}
