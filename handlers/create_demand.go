package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
)

type createDemandDto struct {
	DemandID  int64     `json:"demand_id" gorm:"not null"`
	ShiftID   string    `json:"shift_id" gorm:"not null"`
	UserID    int64     `json:"user_id" gorm:"default:null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
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

func (ss *ShiftService) HandleCreateDemand(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the demand data
	var demand models.Demand
	if err := c.ShouldBindJSON(&demand); err != nil {

		return http.StatusBadRequest, nil, errors.New("invalid shift data")
	}

	// Perform any validation or preprocessing of the shift data if needed

	// Save the demand to the database
	if err := ss.db.Create(&demand).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create shift")
	}
	// Return the created demand as the response
	return http.StatusOK, demand, nil
}
