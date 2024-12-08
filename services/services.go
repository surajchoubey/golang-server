package services

import (
	"api/models"
	"database/sql"
	"errors"
	"log"
)

// GET ALL USERS
func FetchUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query(("SELECT * from users"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GET ONE USER
func FetchUser(db *sql.DB, id *string) (models.User, error) {

	var user models.User
	err := db.QueryRow("SELECT * from users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, sql.ErrNoRows
		}
		return user, err
	}

	return user, nil
}

// CREATE USER
func CreateUser(db *sql.DB, name *string, email *string) (models.User, error) {
	var u models.User
	var id int

	err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", name, email).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return u, nil
}

// UPDATE USER
func UpdateUser(db *sql.DB, id *string, name *string, email *string) (int64, error) {

	result, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", name, email, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New("no rows updated user may not exist")
	}

	return rowsAffected, nil
}

// DELETE USER
func DeleteUser(db *sql.DB, id *string) (int64, error) {

	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New("no rows deleted, user may not exist")
	}

	return rowsAffected, nil
}
