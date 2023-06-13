package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleUpdateDemand godoc
// @Summary update demand by dto
// @Schemes
// @Description update demand by dto
// @Tags demand
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "update demand by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "demand not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /demand/:id [update]
func (ss *ShiftService) HandleUpdateDemand(c *gin.Context) (int, interface{}, error) {

	// update demand id
	demandID := c.Param("id")

	fmt.Println("log1")
	// update dto from req.body
	var dto models.Demand
	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		fmt.Println("log1")
		return http.StatusBadRequest, nil, errors.New("invalid demand dto")
	}

	// update demand from db
	var demand models.Demand

	// check demand exists
	if err := ss.db.Where("demand_id = ?", demandID).First(&demand).Error; err != nil {
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
