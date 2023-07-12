package handlers

import (
	"database/sql"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
)

// HandleGetAllUsers godoc
// @Summary get all users
// @Schemes
// @Description get all users
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "get all users success"
// @Failure 422 {object} RespondJson "users not found"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /users [get]

func (ss *ShiftService) HandleGetAllUsers(c *gin.Context) (int, interface{}, error) {
	var users []models.User
	if err := ss.db.Find(&users).Error; err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, err
		}
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil
}
