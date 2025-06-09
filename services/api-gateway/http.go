package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	var reqBody startTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	response := contracts.APIResponse{Data: "ok"}

	writeJSON(w, http.StatusCreated, response)

	log.Println("SUCCESS")

}

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

}
