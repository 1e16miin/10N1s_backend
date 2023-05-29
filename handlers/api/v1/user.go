package v1

import (
	"net/http"

	"github.com/1e16miin/models"
	"github.com/labstack/echo/v4"
)

func getUsers(c echo.Context) error {
	// Fetch all users from the database
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch users",
		})
	}

	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	// Parse the request body into a User object
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Create the user in the database
	if err := db.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, user)
}
