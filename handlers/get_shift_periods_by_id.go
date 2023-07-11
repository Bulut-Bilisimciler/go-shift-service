package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleGetShiftPeriodByID godoc
// @Summary Get shift_period by id
// @Description Get shift_period by id
// @Tags shift_periods
// @Accept json
// @Produce json
// @Param id path string true "ShiftPeriod ID"
// @Security BearerAuth
// @Success 200 {object} RespondJson "get shift_period by id success"
// @Failure 400 {object} RespondJson "invalid get shift_period by id dto"
// @Failure 422 {object} RespondJson "shift_period not found"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /shift_periods/{id} [get]

func (ss *ShiftService) HandleGetShiftPeriodByID(c *gin.Context) (int, interface{}, error) {

	// Get shift id from req.params
	shift_periodId := c.Param("id")

	// get user
	var shift_period models.ShiftPeriod
	if err := ss.db.Where("shift_period_id = ?", shift_periodId).First(&shift_period).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("shift_period not found")
		} else {
			return http.StatusInternalServerError, nil, errors.New("shift_period not found due to internal error")
		}
	}

	return http.StatusOK, shift_period, nil
}
