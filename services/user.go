package services

import (
	"fmt"
	"purchases-api/db"
	"purchases-api/models"
)

var dbUsers = db.DBUsers

func CreateUser(user models.User) (models.User, error) {
	return dbUsers.Save(user)
}

func GetAllUsers() []models.User {
	return dbUsers.GetAll()
}

func GetUserByID(key string) (interface{}, error) {
	user, error := dbUsers.GetByID(key)
	if error != nil {
		return nil, error
	}
	return user, nil
}

func UpdateUser(key string, user models.User) (interface{}, error) {
	currentUser, err := dbUsers.GetByID(key)
	if err != nil {
		return nil, err
	}
	if user.DNI != currentUser.DNI && user.DNI > 0 {
		currentUser.DNI = user.DNI
	}
	if user.Name != currentUser.Name && user.Name != "" {
		currentUser.Name = user.Name
	}
	if user.LastName != currentUser.LastName && user.LastName != "" {
		currentUser.LastName = user.LastName
	}

	return currentUser, dbUsers.Update(key, currentUser)
}

func DeleteUser(key string) string {
	_, err := dbUsers.GetByID(key)
	if err != nil {
		return fmt.Sprintf("User id: " + key + " doesn't exist")
	}
	err = dbUsers.Delete(key)
	if err != nil {
		return fmt.Sprintf("User id: " + key + " can't be deleted")
	}

	return fmt.Sprintf("User id: " + key + " deleted")
}
