package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valentinrb1/go-api-rest.git/models"
)

var info models.Info

func SubmitProcessingHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Processing request submitted")
}

func GetSummaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
