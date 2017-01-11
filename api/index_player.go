package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/benspotatoes/extrawatch/models"
)

type indexPlayerResp struct {
	Players []*models.Player `json:"players"`
}

const (
	searchParam = "search"
)

func (rtr *Router) indexPlayer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := query.Get(searchParam)

	players, err := rtr.Backend.IndexPlayer(r.Context(), filter)
	if err == sql.ErrNoRows {
		rtr.handleErrorResponse(w, r, http.StatusNotFound, err)
		return
	}

	res := &indexPlayerResp{Players: players}
	blob, err := json.Marshal(res)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
