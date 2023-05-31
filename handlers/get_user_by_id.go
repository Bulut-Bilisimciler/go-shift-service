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
func (ss *ShiftService) HandleGetUserById(c *gin.Context) (int, interface{}, error) {

	// Get user id from req.params
	userId := c.Param("id")

	// get user
	var user models.User
	if err := ss.db.Where("user_id = ?", userId).First(&user).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("user not found")
		} else {
			return http.StatusInternalServerError, nil, errors.New("user not found due to internal error")
		}
	}

	return http.StatusOK, user, nil
}
