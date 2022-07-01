package handlers

import (
	"better-auth/internal/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SessionHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response map[string]interface{}
	var session models.Session
	SID := req.URL.Query().Get("SID")
	ctx := context.Background()
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the response body. "http: request body too large" error.
	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	dec := json.NewDecoder(req.Body)
	// Setup the decoder and call the DisallowUnknownFields() method on it. This will cause Decode() to return a "json: unknown field ..." error
	dec.DisallowUnknownFields()
	switch req.Method {
	case "GET":
		if SID == "" {
			response = models.CreateResponse("failure", "Session ID is Missing", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		var startTime = time.Now()
		res := session.Get(&ctx, SID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Session Get Operation Took: %v Nano Seconds", duration.Nanoseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	case req.Method:
		response = models.CreateResponse("failure", "method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
}
