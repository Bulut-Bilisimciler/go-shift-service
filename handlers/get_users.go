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
	// pagination from req.query
	var params models.Pagination
	if err := c.ShouldBindQuery(&params); !errors.Is(err, nil) {
		return http.StatusBadRequest, nil, errors.New("invalid pagination query: page,limit")
	}

	// get users
	limit := params.Limit
	offset := (params.Page - 1) * params.Limit

	var users []models.User
	if err := ss.db.Limit(limit).Offset(offset).Find(&users).Error; !errors.Is(err, nil) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("users not found")
		}

		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	// return
	return http.StatusOK, users, nil

}
