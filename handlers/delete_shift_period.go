package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleDeleteShiftPeriod handles the delete shift_period request
// @Summary delete shift_period by dto
// @Schemes
// @Description delete shift_period by dto
// @Tags shift_periods
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "delete shift_period success"
// @Failure 400 {object} RespondJson "invalid delete shift_period dto"
// @Failure 422 {object} RespondJson "cannot delete shift_period due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /shift_periods [delete]

func (ss *ShiftService) HandleDeleteShiftPeriod(c *gin.Context) (int, interface{}, error) {
	shiftPeriodId := c.Param("id")

	// delete shift
	var shiftPeriod models.ShiftPeriod
	if err := ss.db.Where("shift_period_id = ?", shiftPeriodId).Delete(&shiftPeriod).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("shift period not found")
		} else {
			return http.StatusInternalServerError, nil, errors.New("shift period not found due to internal error")
		}
	}

	return http.StatusOK, shiftPeriod, nil

}
