package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
)

// HandleCreateShiftPeriod handles the create shift_periods request
// @Summary create shift_periods by dto
// @Schemes
// @Description create shift_periods by dto
// @Tags shift_periods
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "create shift_periods success"
// @Failure 400 {object} RespondJson "invalid create shift_periods dto"
// @Failure 422 {object} RespondJson "cannot create shift_periods due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /shift_periods [post]

func (ss *ShiftService) HandleCreateShiftPeriod(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the shift_periods data
	var shift_period models.ShiftPeriod
	if err := c.ShouldBindJSON(&shift_period); err != nil {

		return http.StatusBadRequest, nil, errors.New("invalid shift period data")
	}

	// Save the shift_periods to the database
	if err := ss.db.Create(&shift_period).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create shift period")
	}

	// Return the created shift_periods as the response
	return http.StatusOK, shift_period, nil
}
