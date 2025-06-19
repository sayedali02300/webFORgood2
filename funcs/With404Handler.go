package groupie

import (
	"net/http"
)

func Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "templates/404.html")
}

func With404Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux, ok := next.(*http.ServeMux)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		_, pattern := mux.Handler(r)
		if pattern == "" {
			Handle404(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
