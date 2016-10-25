package api

import (
	"net/http"

	"github.com/benspotatoes/extrawatch/backend"
	"goji.io"
	"goji.io/pat"
)

type Router struct {
	Backend backend.Backend
}

func NewRouter(b backend.Backend) *goji.Mux {
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

	return mux
}

func (rtr *Router) running(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}