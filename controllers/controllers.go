package controllers

import (
	"api/models"
	"api/services"
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL USERS
func GetUsersController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := services.FetchUsers(db)
		fmt.Println(users, err)
		if err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch users")
		}
		utils.SendSuccessResponse(w, users)
	}
}

// GET USER
func GetUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		user, err := services.FetchUser(db, &id)
		if err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user")
		}
		utils.SendSuccessResponse(w, user)
	}
}

// CREATE USER
func CreateUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User
		json.NewDecoder(r.Body).Decode(&u)

		user, err := services.CreateUser(db, &u.Name, &u.Email)

		if err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		}

		utils.SendSuccessResponse(w, user)
	}
}

// UPDATE USER
func UpdateUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var u models.User
		json.NewDecoder(r.Body).Decode(&u)
		vars := mux.Vars(r)
		id := vars["id"]

		updated, err := services.UpdateUser(db, &id, &u.Name, &u.Email)

		if err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to update user")
		}

		utils.SendSuccessResponse(w, updated)
	}
}

// DELETE USER
func DeleteUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		deleted, err := services.DeleteUser(db, &id)

		if err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to delete user")
		}

		utils.SendSuccessResponse(w, deleted)
	}
}
