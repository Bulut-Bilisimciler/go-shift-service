package handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
)

// HandleGetTodos godoc
// @Summary get todos by pagination
// @Schemes
// @Description get todos by pagination
// @Tags Todos
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination query"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "get todos by pagination success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query: page,limit"
// @Failure 404 {object} handlers.RespondJson "todos not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /todos [get]
func (mss *MySuperService) HandleGetTodos(c *gin.Context) (int, interface{}, error) {
	// pagination from req.query
	var params models.Pagination
	if err := c.ShouldBindQuery(&params); !errors.Is(err, nil) {
		return http.StatusBadRequest, nil, errors.New("invalid pagination query: page,limit")
	}

	/*
		// Get JWT from gin context
		// get jwt user from context
		jwtResponse, ok := c.MustGet("user").(*models.JwtResponse)
		if !ok {
			return http.StatusUnauthorized, nil, errors.New("jwt user is invalid")
		}
		authIdStr := strconv.FormatInt(jwtResponse.AuthId, 10)
	*/

	// get category
	limit := params.Limit
	offset := (params.Page - 1) * params.Limit

	var todos []models.Todo
	if err := mss.db.Limit(limit).Offset(offset).Find(&todos).Error; !errors.Is(err, nil) {
		return http.StatusNotFound, nil, errors.New("todos not found")
	}

	// return
	return http.StatusOK, todos, nil

}
