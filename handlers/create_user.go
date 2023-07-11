package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
)

// HandleCreateUser handles the create user request
// @Summary create user by dto
// @Schemes
// @Description create user by dto
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "create user success"
// @Failure 400 {object} RespondJson "invalid create user dto"
// @Failure 422 {object} RespondJson "cannot create user due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /users [post]

func (ss *ShiftService) HandleCreateUser(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the user data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return http.StatusBadRequest, nil, errors.New("invalid user data")
	}

	// Save the user to the database
	if err := ss.db.Create(&user).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create user")
	}

	// Return the created user as the response
	return http.StatusOK, user, nil
}
