package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// HandleGetDemand godoc
// @Summary get demands by dto
// @Schemes
// @Description get demands by dto
// @Tags demands
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "get demands by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "demands not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /demands [get]

func (ss *ShiftService) HandleGetDemand(c *gin.Context) (int, interface{}, error) {

	// get demands
	var demands []models.Demand
	if err := ss.db.Find(&demands).Error; err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, err
		}
		log.Println(err)
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, demands, nil
}
