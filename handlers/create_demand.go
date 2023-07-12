package handlers

import (
	"buluttan/shift-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleCreateDemand handles the create demand request
// @Summary create demand by dto
// @Schemes
// @Description create demand by dto
// @Tags demands
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "create demand success"
// @Failure 400 {object} RespondJson "invalid create demand dto"
// @Failure 422 {object} RespondJson "cannot create demand due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /demands [post]

func (ss *ShiftService) HandleCreateDemand(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the demand data
	var demand models.Demand
	if err := c.ShouldBindJSON(&demand); err != nil {
		return http.StatusBadRequest, nil, errors.New("invalid demand data")
	}

	// Save the demand to the database
	if err := ss.db.Create(&demand).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create demand")
	}

	// Return the created demand as the response
	return http.StatusOK, demand, nil
}
