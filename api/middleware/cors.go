package middleware

import "net/http"

func Cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header()["Access-Control-Allow-Origin"] = []string{"*"}
		handler.ServeHTTP(w, r)
	})
}
