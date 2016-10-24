package models

var (
	validModes map[string]bool
	modes      = []string{
		// Escort, Assault, Hybrid
		"attack",
		"defense",
		// LijiangTower
		"market",
		"gardens",
		"center",
		// Nepal
		"village",
		"shrine",
		"sanctum",
		// Ilios
		"lighthouse",
		"well",
		"ruins",
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
	Count   int       `json:"count"`
	Players []*Player `json:"players"`
	Mode    string    `json:"mode"`
	Result  *Result   `json:"result"`
	Notes   string    `json:"notes"`
}

func (r *Round) validate() error {
	// Count
	if r.Count < 0 {
		return errInvalidRoundCount
	}

	// Players
	if len(r.Players) != playerCount {
		return errInvalidRoundPlayers
	}
	for _, player := range r.Players {
		// Player
		if playErr := player.Validate(); playErr != nil {
			return playErr
		}
	}

	// Mode
	// TODO - Validate mode with map context
	if !validModes[r.Mode] {
		return errInvalidMode
	}

	// Result
	if resErr := r.Result.validate(); resErr != nil {
		return resErr
	}

	// Notes
	if len(r.Notes) > 500 {
		return errInvalidNotes
	}

	return nil
}

type Result struct {
	TimeLeft    int `json:"time_left"`
	PercentDiff int `json:"percent_diff"`
	PointsTaken int `json:"points_taken"`
}

func (m *Result) validate() error {
	switch {
	case m.TimeLeft != 0 && m.PercentDiff == 0 && m.PointsTaken == 0:
		// Time remaining means the payload was pushed to the last point
		return nil
	case m.PercentDiff != 0 && m.TimeLeft == 0 && m.PointsTaken == 0:
		// Positive percent diff means a win, negative means a loss
		return nil
	case m.PointsTaken >= 0 && m.TimeLeft == 0 && m.PercentDiff == 0:
		// No points taken with no time left means either the attack was a super
		// disaster or the defense was a great success
		return nil
	default:
		return errInvalidRoundResult
	}
}
