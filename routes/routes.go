package routes

import (
	"api/controllers"
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/users", controllers.GetUsersController(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserController(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUserController(db)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUserController(db)).Methods("PATCH")
	router.HandleFunc("/users/{id}", controllers.DeleteUserController(db)).Methods("DELETE")
}
