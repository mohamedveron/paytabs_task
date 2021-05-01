package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := s.svc.GetAccounts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
