package api

import (
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) deleteMatch(w http.ResponseWriter, r *http.Request) {
	matchID := pat.Param(r, "match_id")

	err := rtr.Backend.DeleteMatch(r.Context(), matchID)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusNoContent, nil)
}
