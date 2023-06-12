package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
)

type createShiftDto struct {
	ShiftID   string    `json:"shift_id" gorm:"not null"`
	UserID    int64     `json:"user_id" gorm:"default:null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
	MadeField string    `json:"made_field" gorm:"default:null"`
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

func (ss *ShiftService) HandleCreateShift(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the shift data
	var shift models.Shift
	if err := c.ShouldBindJSON(&shift); err != nil {

		return http.StatusBadRequest, nil, errors.New("invalid shift data")
	}

	// Perform any validation or preprocessing of the shift data if needed

	// Save the shift to the database
	if err := ss.db.Create(&shift).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create shift")
	}
	// Return the created shift as the response
	return http.StatusOK, shift, nil
}
