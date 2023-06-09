package handlers

import (
	"errors"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleUpdateUser godoc
// @Summary update users by dto
// @Schemes
// @Description update users by dto
// @Tags users
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "update users by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "users not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /users/:id [update]

func (ss *ShiftService) HandleUpdateUser(c *gin.Context) (int, interface{}, error) {

	// get user id
	employeeID := c.Param("id")

	// get dto from req.body
	var dto models.User
	if err := c.ShouldBindJSON(&dto); err != nil {
		return http.StatusBadRequest, nil, errors.New("invalid user dto")
	}

	// get user from db
	var user models.User

	// check user exists
	if err := ss.db.Where("employee_id = ?", employeeID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusUnprocessableEntity, nil, errors.New("user not found")
		}
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	// update user
	if err := ss.db.Model(&user).Updates(dto).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	return http.StatusOK, user, nil
}
