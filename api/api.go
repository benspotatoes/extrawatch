package api

import (
	"net/http"

	"github.com/benspotatoes/extrawatch/api/middleware"
	"github.com/benspotatoes/extrawatch/backend"
	"goji.io"
	"goji.io/pat"
)

type Config struct {
	Cors bool
}

type Router struct {
	Backend backend.Backend
}

func NewRouter(b backend.Backend, c *Config) *goji.Mux {
	router := &Router{
		Backend: b,
	}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/running"), router.running)

	mux.HandleFunc(pat.Get("/"), router.indexMatch)
	mux.HandleFunc(pat.Get("/match/:match_id"), router.getMatch)
	mux.HandleFunc(pat.Put("/match/:match_id"), router.updateMatch)
	mux.HandleFunc(pat.Delete("/match/:match_id"), router.deleteMatch)
	mux.HandleFunc(pat.Post("/match"), router.postMatch)

	mux.HandleFunc(pat.Get("/round/:round_id"), router.getRound)
	mux.HandleFunc(pat.Put("/round/:round_id"), router.updateRound)
	mux.HandleFunc(pat.Delete("/round/:round_id"), router.deleteRound)
	mux.HandleFunc(pat.Post("/round"), router.postRound)

	mux.HandleFunc(pat.Get("/player"), router.indexPlayer)
	mux.HandleFunc(pat.Get("/player/:player_id"), router.getPlayer)
	mux.HandleFunc(pat.Put("/player/:player_id"), router.updatePlayer)
	mux.HandleFunc(pat.Delete("/player/:player_id"), router.deletePlayer)
	mux.HandleFunc(pat.Post("/player"), router.postPlayer)

	mux.HandleFunc(pat.Get("/player/:player_id/round/:round_id"), router.getPlayerRound)
	mux.HandleFunc(pat.Delete("/player/:player_id/round/:round_id"), router.deletePlayerRound)
	mux.HandleFunc(pat.Post("/player/round"), router.postPlayerRound)

	if c.Cors {
		mux.Use(middleware.Cors)
	}

	return mux
}

func (rtr *Router) running(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
