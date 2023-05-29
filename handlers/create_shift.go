package handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
)

type createShiftDto struct {
	UserName string `form:"username" json:"user_name" binding:"omitempty"`
	UserId   int64  `form:"user_id,default=1" json:"user_id" binding:"omitempty,min=1,max=64"`
	Shift    string `form:"shift" json:"shift"`
	Status   int    `form:"status,default=1" json:"status,omitempty" binding:"omitempty,min=1,max=3"`
}

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

func (ss *ShiftService) HandleCreateShift(c *gin.Context) (int, interface{}, error) {
	var params createShiftDto
	if err := c.ShouldBindJSON(&params); !errors.Is(err, nil) {
		return http.StatusBadRequest, nil, errors.New("invalid create shift dto")
	}

	// map req to model
	var entity models.Shift
	// entity.Shift = params.Shift
	// entity.User.UserId = params.UserId

	// Create shift
	if err := ss.db.Create(&entity).Error; !errors.Is(err, nil) {
		return http.StatusUnprocessableEntity, nil, errors.New("cannot create shift due to invalid value while saving db")
	}

	return http.StatusOK, entity, nil
}
