package models

var (
	validHeroes map[string]bool
	heroes      = []string{
		"ana",
		"bastion",
		"dva",
		"genji",
		"hanzo",
		"junkrat",
		"lucio",
		"mcCree",
		"mei",
		"mercy",
		"pharah",
		"reaper",
		"reinhardt",
		"roadhog",
		"soldier76",
		"symmetra",
		"torbjorn",
		"tracer",
		"widowmaker",
		"winston",
		"zarya",
		"zenyatta",
	}
)

const (
	rando = "*"
)

func init() {
	validHeroes = make(map[string]bool, len(heroes))
	for _, name := range heroes {
		validHeroes[name] = true
	}
}

type Player struct {
	Name string `json:"name"`
	// TODO - Should we enum this bad boy?
	// I want to make this a write-in field so we don't have to worry about a
	// dropdown selector (to handle enums) somehow conveying multiple-hero
	// seletions
	// TODO - Should we limit this to one selection?
	Heroes []string `json:"heroes"`
}

func (p *Player) Validate() error {
	// Name
	if p.Name == "" {
		// Player must have either a specified name or be a pick-up
		p.Name = rando
	}

	// Heroes
	if len(p.Heroes) > 3 {
		return errInvalidHeroes
	}
	for _, hero := range p.Heroes {
		// Hero
		if !validHeroes[hero] {
			return errInvalidHero
		}
	}

	return nil
}
