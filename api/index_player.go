package api

import "net/http"

func (rtr *Router) indexPlayer(w http.ResponseWriter, r *http.Request) {
	rtr.handleErrorResponse(w, r, http.StatusNotImplemented, nil)
}
