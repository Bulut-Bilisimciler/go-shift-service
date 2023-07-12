package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleDeleteDemand handles the delete demand request
// @Summary delete demand by dto
// @Schemes
// @Description delete demand by dto
// @Tags demands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "delete demand success"
// @Failure 400 {object} RespondJson "invalid delete demand dto"
// @Failure 422 {object} RespondJson "cannot delete demand due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /demands [delete]

func (ss *ShiftService) HandleDeleteDemands(c *gin.Context) (int, interface{}, error) {
	demandId := c.Param("id")

	// delete demand
	var demand models.Demand
	if err := ss.db.Where("demand_id = ?", demandId).Delete(&demand).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("demand not found")
		} else {
			return http.StatusInternalServerError, nil, errors.New("demand not found due to internal error")
		}
	}

	return http.StatusOK, demand, nil

}
