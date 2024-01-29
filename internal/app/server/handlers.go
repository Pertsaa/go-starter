package server

import (
	"net/http"
)

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respondJSON(w, r, "Hello, World!", http.StatusOK)
	}
}
