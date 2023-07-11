package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleGetDemandById handles the get demand by id request
// @Summary get demand by id
// @Schemes
// @Description get demand by id
// @Tags demands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "get demand by id success"
// @Failure 400 {object} RespondJson "invalid get demand by id dto"
// @Failure 422 {object} RespondJson "cannot get demand by id due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /demands/{id} [get]

func (ss *ShiftService) HandleGetDemandById(c *gin.Context) (int, interface{}, error) {

	// Get shift id from req.params
	demandId := c.Param("id")

	// get demand
	var demand models.Demand
	if err := ss.db.Where("demand_id = ?", demandId).First(&demand).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("demand not found")
		} else {
			return http.StatusInternalServerError, nil, errors.New("demand not found due to internal error")
		}
	}

	return http.StatusOK, demand, nil
}
