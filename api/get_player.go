package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) getPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := pat.Param(r, "player_id")

	player, err := rtr.Backend.SelectPlayer(r.Context(), playerID)
	if err == sql.ErrNoRows {
		rtr.handleErrorResponse(w, r, http.StatusNotFound, nil)
		return
	} else if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	blob, err := json.Marshal(player)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
