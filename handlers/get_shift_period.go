package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// HandleGetShiftPeriod godoc
// @Summary get shift_periods by dto
// @Schemes
// @Description get shift_periods by dto
// @Tags shift_periods
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "get shift_periods by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "shift_periods not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /shift_periods [get]

func (ss *ShiftService) HandleGetShiftPeriod(c *gin.Context) (int, interface{}, error) {
	var shift_period []models.ShiftPeriod
	if err := ss.db.Find(&shift_period).Error; err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, err
		}
		log.Println(err)
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, shift_period, nil
}
