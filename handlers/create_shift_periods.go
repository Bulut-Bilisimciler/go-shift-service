package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
)

type createShiftPeriodDto struct {
	ShiftPeriodID int64     `json:"shift_period_id" gorm:"not null"`
	StartTime     time.Time `json:"start_time" gorm:"not null"`
	EndTime       time.Time `json:"end_time" gorm:"not null"`
}

// HandleCreateShift godoc
// @Summary create shift by dto
// @Schemes
// @Description create shift by dto
// @Tags shifts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "create shift success"
// @Failure 400 {object} handlers.RespondJson "invalid create shift dto"
// @Failure 422 {object} handlers.RespondJson "cannot create shift due to db error"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /shifts [post]

func (ss *ShiftService) HandleCreateShiftPeriod(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the shift_periods data
	var shift_period models.ShiftPeriod
	if err := c.ShouldBindJSON(&shift_period); err != nil {

		return http.StatusBadRequest, nil, errors.New("invalid shift_period data")
	}

	// Perform any validation or preprocessing of the shift data if needed

	// Save the shift_periods to the database
	if err := ss.db.Create(&shift_period).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create shift_period")
	}
	// Return the created shift_periods as the response
	return http.StatusOK, shift_period, nil
}
