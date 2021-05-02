package api

import (
	"encoding/json"
	"net/http"
)

func (s * Server) MakeTransfer(w http.ResponseWriter, r *http.Request) {

	var transferRequest Transfer
	if err := json.NewDecoder(r.Body).Decode(&transferRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.svc.MakeTransfer(transferRequest.From, transferRequest.To, transferRequest.Balance)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Transfer Succeeded")

}
