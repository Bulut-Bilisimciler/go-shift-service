package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleGetShiftById handles the get shift by id request
// @Summary get shift by id
// @Schemes
// @Description get shift by id
// @Tags shifts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "get shift by id success"
// @Failure 400 {object} RespondJson "invalid get shift by id dto"
// @Failure 422 {object} RespondJson "cannot get shift by id due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /shifts/{id} [get]

func (ss *ShiftService) HandleGetShiftById(c *gin.Context) (int, interface{}, error) {

	// Get shift id from req.params
	shiftId := c.Param("id")

	// get shift
	var shift models.Shift
	if err := ss.db.Where("shift_id = ?", shiftId).First(&shift).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("shift not found")
		} else {
			return http.StatusInternalServerError, nil, errors.New("shift not found due to internal error")
		}
	}

	return http.StatusOK, shift, nil
}
