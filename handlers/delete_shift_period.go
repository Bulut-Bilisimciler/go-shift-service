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
// @Router /shifts [get]
func (ss *ShiftService) HandleDeleteShiftPeriod(c *gin.Context) (int, interface{}, error) {

	// pagination from req.query
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
