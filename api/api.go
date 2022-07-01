package api

import (
	"better-auth/api/handlers"
	"better-auth/configs"
	"log"
	"net/http"
)

/**
API ENDPOINTS
CRUD ANY MAP[STRING] INTERFACE IF KEY SET TO TRUE
CRUD USER
*/

func GetMuxAPI() *http.ServeMux {
	log.Print("Initializing Rest Endpoints " + configs.Port)
	mux := http.NewServeMux()
	/**
	CREATE		USER OBJECT in DB
	RETRIEVE	USER OBJECT from DB
	UPDATE		USER OBJECT in DB
	DELETE		USER OBJECT in DB
	*/
	mux.HandleFunc("/api/v1/user/new", handlers.CreateUserHandler)
	mux.HandleFunc("/api/v1/user", handlers.GetUserHandler)
	//mux.HandleFunc("/api/v1/user/all", handlers.GetUsersHandler)
	mux.HandleFunc("/api/v1/user/update", handlers.UpdateUserHandler)
	mux.HandleFunc("/api/v1/user/delete", handlers.DeleteUserHandler)

	mux.HandleFunc("/api/v1/user/login", handlers.LoginHandler)
	mux.HandleFunc("/api/v1/user/logout", handlers.LogoutHandler)

	mux.HandleFunc("/api/v1/user/session", handlers.SessionHandler)
	return mux
}
