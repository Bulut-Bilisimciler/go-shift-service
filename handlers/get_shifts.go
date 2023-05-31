package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
func (ss *ShiftService) HandleGetShift(c *gin.Context) (int, interface{}, error) {

	// get shifts
	var shifts []models.Shift
	if err := ss.db.Find(&shifts).Error; err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, err
		}
		log.Println(err)
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, shifts, nil
}
