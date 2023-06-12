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
func (ss *ShiftService) HandleUpdateDemand(c *gin.Context) (int, interface{}, error) {

	// get demand id
	demandID := c.Param("id")

	// get dto from req.body
	var dto models.Demand
	if err := c.ShouldBindJSON(&dto); err != nil {
		return http.StatusBadRequest, nil, errors.New("invalid demand dto")
	}

	// get demand from db
	var demand models.Demand

	// check demand exists
	if err := ss.db.Where("id = ?", demandID).First(&demand).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusUnprocessableEntity, nil, errors.New("demand not found")
		}
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	// update demand
	if err := ss.db.Model(&demand).Updates(dto).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	return http.StatusOK, demand, nil
}
