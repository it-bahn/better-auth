package handlers

import (
	"better-auth/internal/models"
	"better-auth/internal/models/sessions"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	log.Printf("AuthMiddleware called")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization")
		}
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		var response map[string]interface{}
		var session sessions.Session
		ctx := context.Background()
		response = session.FindOne(&ctx, r.Header.Get("Authorization"))
		// Check if the user is authenticated
		if response["status"] == "failure" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func FilteredMiddleware(next http.Handler) http.Handler {
	log.Printf("FilteredMiddleware called")

	var response map[string]interface{}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization")
		}
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Content-Type", "application/json")
		var userID = req.URL.Query().Get("id")
		if req.Method == "GET" || req.Method == "DELETE" || req.Method == "PUT" {
			if userID == "" {
				log.Printf("Error: No user ID provided")
				response = models.CreateResponse("failure", "No user ID provided", http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}
		}
		// Call the next handler
		next.ServeHTTP(w, req)
	})
}
