package server

import (
	"encoding/json"
	"net/http"
)

func (s *server) respondJSON(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.respondJSON(w, r, err, http.StatusBadRequest)
		}
	}
}
