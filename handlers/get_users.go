package handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleGetShifts godoc
// @Summary get users by dto
// @Schemes
// @Description get users by dto
// @Tags users
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "get shifts by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "shifts not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /users [get]
func (ss *ShiftService) HandleGetUsers(c *gin.Context) (int, interface{}, error) {

	// get users
	var users []models.User
	err := ss.db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusOK, users, errors.New("no users found")
		}

		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil

}
