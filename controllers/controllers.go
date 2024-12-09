package controllers

import (
	"api/models"
	"api/services"
	"api/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GET ALL USERS
func GetUsersController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := services.FetchUsers(db)
		if err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch users")
		}
		utils.SendSuccessResponse(w, users)
	}
}

// GET USER
func GetUserController(db *gorm.DB) http.HandlerFunc {
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
func CreateUserController(db *gorm.DB) http.HandlerFunc {
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
func UpdateUserController(db *gorm.DB) http.HandlerFunc {
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
func DeleteUserController(db *gorm.DB) http.HandlerFunc {
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
