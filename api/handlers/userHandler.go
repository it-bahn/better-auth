package handlers

import (
	"better-auth/internal/models"
	"better-auth/internal/models/users"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/**
 * Create a new Object
 */
func UserHandler(w http.ResponseWriter, req *http.Request) {
	var response map[string]interface{}
	var incoming users.User
	var userID = req.URL.Query().Get("id")
	ctx := context.Background()
	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the response body. "http: request body too large" error.
	dec := json.NewDecoder(req.Body)
	// Setup the decoder and call the DisallowUnknownFields() method on it. This will cause Decode() to return a "json: unknown field ..." error
	dec.DisallowUnknownFields()
	switch req.Method {
	case "POST":
		// Decode the request body into the struct.
		err := dec.Decode(&incoming)
		if err != nil {
			response = GetErrorFromRequestBody(w, err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		response = ValidateReqBody(w, incoming)
		if response["status"] == "failure" {
			json.NewEncoder(w).Encode(response)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var startTime = time.Now()
		response = incoming.Create(&ctx)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Insert Operation Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	case "GET":
		var startTime = time.Now()
		response = incoming.FindOne(&ctx, userID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Find One Operation Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	case "PUT":
		err := dec.Decode(&incoming)
		if err != nil {
			response = GetErrorFromRequestBody(w, err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		response = ValidateReqBody(w, incoming)
		if response["status"] == "failure" {
			json.NewEncoder(w).Encode(response)
			return
		}
		var startTime = time.Now()
		response = incoming.Update(&ctx, userID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Find and Update One Operation Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	case "DELETE":
		var startTime = time.Now()
		response = incoming.Delete(&ctx, userID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Delete One Operation Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case req.Method:
		response = models.CreateResponse("failure", "method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

}
