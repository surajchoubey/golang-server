package services

import (
	"api/models"

	"gorm.io/gorm"
)

// GET ALL USERS
func FetchUsers(db *gorm.DB) ([]models.User, error) {

	var users []models.User
	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// GET ONE USER
func FetchUser(db *gorm.DB, id *string) (models.User, error) {

	var user models.User

	result := db.First(&user, id)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// CREATE USER
func CreateUser(db *gorm.DB, name *string, email *string) (models.User, error) {

	user := models.User{
		Name:  *name,
		Email: *email,
	}

	result := db.Create(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// UPDATE USER
func UpdateUser(db *gorm.DB, id *string, name *string, email *string) (models.User, error) {

	var user models.User
	result := db.Model(&user).Where("id = ?", id).Updates(models.User{Name: *name, Email: *email})

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// DELETE USER
func DeleteUser(db *gorm.DB, id *string) (models.User, error) {

	var user models.User
	result := db.Delete(&user, id)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
