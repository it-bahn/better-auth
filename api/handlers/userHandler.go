package handlers

import (
	"better-auth/internal/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/**
 * Create a new user
 */
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response map[string]interface{}
	var requestData models.User

	Key := req.URL.Query().Get("key")
	Value := req.URL.Query().Get("value")
	log.Printf("key: %s, value: %s", Key, Value)

	ctx := context.Background()
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the response body. "http: request body too large" error.
	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	dec := json.NewDecoder(req.Body)
	// Setup the decoder and call the DisallowUnknownFields() method on it. This will cause Decode() to return a "json: unknown field ..." error
	dec.DisallowUnknownFields()
	switch req.Method {
	case "POST":
		// Decode the request body into the struct.
		err := dec.Decode(&requestData)
		if err != nil {
			response = ParseErrorFromRequest(w, response, err)
			json.NewEncoder(w).Encode(response)
			return
		}
		if requestData.IsEmptyMandatory() {
			response = models.CreateResponse("failure", "Request body is missing mandatory fields", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		var startTime = time.Now()
		res := requestData.Create(&ctx)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Insert Operation Took: %v Nano Seconds", duration.Nanoseconds())
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

/**
 * Update a new user
 */
func UpdateUserHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response map[string]interface{}
	var requestData models.User

	userID := req.URL.Query().Get("userID")
	log.Printf("userID: %s", userID)
	if userID == "" {
		response = models.CreateResponse("failure", "userID is missing", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	ctx := context.Background()
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the response body. "http: request body too large" error.
	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	dec := json.NewDecoder(req.Body)
	// Setup the decoder and call the DisallowUnknownFields() method on it. This will cause Decode() to return a "json: unknown field ..." error
	dec.DisallowUnknownFields()
	switch req.Method {
	case "PUT":
		// Decode the request body into the struct.
		err := dec.Decode(&requestData)
		if err != nil {
			response = ParseErrorFromRequest(w, response, err)
			json.NewEncoder(w).Encode(response)
			return
		}
		if requestData.IsEmptyEntirely() {
			response = models.CreateResponse("failure", "Request body is missing all fields", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		var startTime = time.Now()
		res := requestData.Update(&ctx, userID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Update Operation Took: %v Nano Seconds", duration.Nanoseconds())
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

/**
 * Get a new user
 */
func GetUserHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response map[string]interface{}
	var requestData models.User

	userID := req.URL.Query().Get("userID")
	log.Printf("userID: %s", userID)
	if userID == "" {
		response = models.CreateResponse("failure", "userID is missing", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	ctx := context.Background()
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the response body. "http: request body too large" error.
	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	dec := json.NewDecoder(req.Body)
	// Setup the decoder and call the DisallowUnknownFields() method on it. This will cause Decode() to return a "json: unknown field ..." error
	dec.DisallowUnknownFields()
	switch req.Method {
	case "GET":
		var startTime = time.Now()
		res := requestData.FindOne(&ctx, userID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Find Operation Took: %v Nano Seconds", duration.Nanoseconds())
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

/**
 * Delete A User
 */
func DeleteUserHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response map[string]interface{}
	var requestData models.User

	userID := req.URL.Query().Get("userID")
	log.Printf("userID: %s", userID)
	if userID == "" {
		response = models.CreateResponse("failure", "userID is missing", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	ctx := context.Background()
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the response body. "http: request body too large" error.
	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	dec := json.NewDecoder(req.Body)
	// Setup the decoder and call the DisallowUnknownFields() method on it. This will cause Decode() to return a "json: unknown field ..." error
	dec.DisallowUnknownFields()
	switch req.Method {
	case "DELETE":
		var startTime = time.Now()
		res := requestData.Delete(&ctx, userID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Delete Operation Took: %v Nano Seconds", duration.Nanoseconds())
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
