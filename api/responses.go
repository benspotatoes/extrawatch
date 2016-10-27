package api

import (
	"fmt"
	"log"
	"net/http"
)

func (rtr *Router) handleSuccessResponse(w http.ResponseWriter, r *http.Request, status int, body interface{}) {
	log.Printf("[%s %s] returning %d", r.Method, r.RequestURI, status)
	w.WriteHeader(status)
	if body != nil {
		_, err := w.Write(body.([]byte))
		if err != nil {
			log.Printf("failed to write body %q", err.Error())
		}
	}
}

func (rtr *Router) handleErrorResponse(w http.ResponseWriter, r *http.Request, status int, err error) {
	var msg string
	if err != nil {
		msg = fmt.Sprintf("with error message %q", err.Error())
	}
	log.Printf("[%s %s] returning %d %s", r.Method, r.RequestURI, status, msg)
	w.WriteHeader(status)
}
