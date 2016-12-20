package api

import (
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) deleteRound(w http.ResponseWriter, r *http.Request) {
	roundID := pat.Param(r, "round_id")

	err := rtr.Backend.DeleteRound(r.Context(), roundID)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusNoContent, nil)
}
