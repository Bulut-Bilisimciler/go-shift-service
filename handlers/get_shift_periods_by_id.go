package handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleGetShifts godoc
// @Summary get shifts by dto
// @Schemes
// @Description get shifts by dto
// @Tags shifts
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "get shifts by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "shifts not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /users [get]
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
