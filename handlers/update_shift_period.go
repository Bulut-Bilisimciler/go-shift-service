package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleUpdateShiftPeriod godoc
// @Summary update shift_periods by dto
// @Schemes
// @Description update shift_periods by dto
// @Tags shift_periods
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "update shift_periods by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "shift_periods not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /shift_periods/:id [update]

func (ss *ShiftService) HandleUpdateShiftPeriod(c *gin.Context) (int, interface{}, error) {

	// update shift_period id
	shiftPeriodID := c.Param("id")

	// update dto from req.body
	var dto models.ShiftPeriod

	fmt.Println("log1")
	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		fmt.Println("log1")
		return http.StatusBadRequest, nil, errors.New("invalid shift_period dto")
	}

	// update shift_period from db
	var shift_period models.ShiftPeriod

	// check shift_period exists
	if err := ss.db.Where("shift_period_id = ?", shiftPeriodID).First(&shift_period).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusUnprocessableEntity, nil, errors.New("shift_period not found")
		}
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	// update shift_period
	if err := ss.db.Model(&shift_period).Updates(dto).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	return http.StatusOK, shift_period, nil
}
