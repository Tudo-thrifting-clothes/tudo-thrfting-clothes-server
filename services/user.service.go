package services

import (
	"errors"
	"tudo-thrifting-server/models" // Use the correct local module path
)

// Dummy data (usually this would come from a database)
var users = []models.User{
	{ID: "1", Name: "Alice", Age: 30},
	{ID: "2", Name: "Bob", Age: 25},
}

// GetUsers returns a list of users
func GetUsers() ([]models.User, error) {
	// In a real application, you would query the database here.
	return users, nil
}

// GetUserByID returns a user by their ID
func GetUserByID(id string) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
