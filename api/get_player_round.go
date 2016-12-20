package api

import (
	"encoding/json"
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) getPlayerRound(w http.ResponseWriter, r *http.Request) {
	playerID := pat.Param(r, "player_id")
	roundID := pat.Param(r, "round_id")

	playerRound, err := rtr.Backend.SelectPlayerRound(r.Context(), playerID, roundID)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if len(playerRound.Heroes) < 1 {
		rtr.handleErrorResponse(w, r, http.StatusNotFound, nil)
		return
	}

	blob, err := json.Marshal(playerRound)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
