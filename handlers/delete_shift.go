package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleDeleteShift handles the delete shift request
// @Summary delete shift by dto
// @Schemes
// @Description delete shift by dto
// @Tags shifts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "delete shift success"
// @Failure 400 {object} RespondJson "invalid delete shift dto"
// @Failure 422 {object} RespondJson "cannot delete shift due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /shifts [delete]

func (ss *ShiftService) HandleDeleteShift(c *gin.Context) (int, interface{}, error) {
	shiftId := c.Param("id")

	// delete shift
	var shift models.Shift
	if err := ss.db.Where("shift_id = ?", shiftId).Delete(&shift).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("shift not found")
		} else {
			return http.StatusInternalServerError, nil, errors.New("shift not found due to internal error")
		}
	}

	return http.StatusOK, shift, nil

}
