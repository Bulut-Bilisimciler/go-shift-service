package handlers

import (
	"buluttan/shift-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

var fields = []string{
	"BAO_Sabit",
	"BAO_Mobil_GB",
	"BAO_IaaS_Otomasyon",
	"BAO_PaaS",
	"BAO_OSS",
	"Windows_Sabit_Mobil_GB",
	"Depolama_Sabit",
	"Depolama_Mobil_GB",
	"Yedekleme_Sabit",
	"Yedekleme_Mobil_GB",
	"Musteri_Depolama_Yedekleme",
	"Musteri_L1_Destek",
}

func (ss *ShiftService) HandleCreateShift(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the shift data
	var shift models.Shift
	if err := c.ShouldBindJSON(&shift); err != nil {
		return http.StatusBadRequest, nil, errors.New("invalid shift data")
	}

	// Save the shift to the database
	if err := ss.db.Create(&shift).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create shift")
	}

	// Return the created shift as the response
	return http.StatusOK, shift, nil
}
