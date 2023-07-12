package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleDeleteUser handles the delete user request
// @Summary delete user by dto
// @Schemes
// @Description delete user by dto
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "delete user success"
// @Failure 400 {object} RespondJson "invalid delete user dto"
// @Failure 422 {object} RespondJson "cannot delete user due to db error"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /users [delete]

func (ss *ShiftService) HandleDeleteUser(c *gin.Context) (int, interface{}, error) {

	// get user id
	employeeID := c.Param("id")

	// get user from db
	var user models.User
	// delete user
	if err := ss.db.Where("id = ?", employeeID).Delete(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("user not found")
		}
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	return http.StatusOK, user, nil
}
