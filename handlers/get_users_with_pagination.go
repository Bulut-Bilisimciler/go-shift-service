package handlers

import (
	"database/sql"
	"net/http"

	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
)

// HandleGetUsersWithPagination godoc
// @Summary get users with pagination
// @Schemes
// @Description get users with pagination
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} RespondJson "get users with pagination success"
// @Failure 400 {object} RespondJson "invalid pagination query"
// @Failure 422 {object} RespondJson "users not found"
// @Failure 500 {object} RespondJson "internal server error"
// @Router /users [get]

func (ss *ShiftService) HandleGetUsersWithPagination(c *gin.Context) (int, interface{}, error) {
	var params models.Pagination
	if err := c.ShouldBindQuery(&params); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var users []models.User
	if err := ss.db.Limit(params.Limit).Offset((params.Page - 1) * params.Limit).Find(&users).Error; err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, err
		}
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil
}

/*
func (ss *ShiftService) HandleGetUsers(c *gin.Context) (int, interface{}, error) {
	// pagination from req.query
	var params models.Pagination
	if err := c.ShouldBindQuery(&params); !errors.Is(err, nil) {
		return http.StatusBadRequest, nil, errors.New("invalid pagination query: page,limit")
	}

	// get users from db
	var users []models.User

	// if pagination not exists send all users
	if params.Page == 0 || params.Limit == 0 {
		if err := ss.db.Find(&users).Error; err != nil {
			if err == sql.ErrNoRows {
				return http.StatusNotFound, nil, err
			}
			return http.StatusInternalServerError, nil, err
		}
	} else {
		if err := ss.db.Limit(params.Limit).Offset((params.Page - 1) * params.Limit).Find(&users).Error; err != nil {
			if err == sql.ErrNoRows {
				return http.StatusNotFound, nil, err
			}
			return http.StatusInternalServerError, nil, err
		}
	}

	// return users with pagination
	return http.StatusOK, users, nil
}*/
