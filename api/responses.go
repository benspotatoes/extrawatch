package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type resp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

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
	if err != nil {
		jsonErr := json.NewEncoder(w).Encode(&resp{StatusCode: status, Message: err.Error()})
		if jsonErr != nil {
			log.Printf("failed to write body %q", jsonErr.Error())
		}
	}
}
