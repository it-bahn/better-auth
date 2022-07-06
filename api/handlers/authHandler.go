package handlers

import (
	"better-auth/internal/models"
	"better-auth/internal/models/auth"
	"better-auth/internal/models/sessions"
	"better-auth/internal/models/users"
	"better-auth/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	var response map[string]interface{}
	var incoming users.User
	ctx := context.Background()
	dec := json.NewDecoder(req.Body)
	switch req.Method {
	case "POST":
		// Decode the request body into the struct.
		_ = dec.Decode(&incoming)
		var startTime = time.Now()
		response = incoming.Create(&ctx)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Insert Operation Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	case req.Method:
		response = models.CreateResponse("failure", "method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var response map[string]interface{}
	var authE auth.AuthEmail
	ctx := context.Background()
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the response body. "http: request body too large" error.
	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	dec := json.NewDecoder(req.Body)
	// Setup the decoder and call the DisallowUnknownFields() method on it. This will cause Decode() to return a "json: unknown field ..." error
	//dec.DisallowUnknownFields()
	switch req.Method {
	case "POST":
		// Decode the request body into the struct.
		err := dec.Decode(&authE)
		if err != nil {
			response = GetErrorFromRequestBody(w, err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		status := utils.IsValidEmail(authE.Email)
		if status == false && authE.Email != "" {
			response := models.CreateResponse("failure", "Invalid Email", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		log.Printf("LoginHandler: %v", authE)
		var startTime = time.Now()
		response := authE.Login(&ctx)
		if response["status"] == "success" {
			var Session sessions.Session
			resp := Session.Create(&ctx, authE.Email, response["user_id"].(string), authE)
			response = map[string]interface{}{
				"status_inserted": resp["_id"],
				"user_id":         response["user_id"],
				"session_id":      resp["session_id"],
				"status":          "success",
				"message":         "login successful",
			}
		}
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Login Operation Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	case req.Method:
		response = models.CreateResponse("failure", "method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var response map[string]interface{}
	var sessionID = req.URL.Query().Get("id")
	var session sessions.Session
	ctx := context.Background()
	switch req.Method {
	case "POST":
		var startTime = time.Now()
		response := session.Delete(&ctx, sessionID)
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Logout Operation Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	case req.Method:
		response = models.CreateResponse("failure", "method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
}
