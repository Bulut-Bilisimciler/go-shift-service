package handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
)

type createUserDto struct {
	UserID     int64  `json:"user_id"  gorm:"default:null"`
	Name       string `json:"name"  gorm:"default:null"`
	Surname    string `json:"surname"  gorm:"default:null"`
	Email      string `json:"email"  gorm:"default:null"`
	Made_Field string `json:"made_field"  gorm:"default:null"`
	Nickname   string `json:"nickname"  gorm:"default:null"`
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

func (ss *ShiftService) HandleCreateUser(c *gin.Context) (int, interface{}, error) {
	// Parse the request body to get the user data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {

		return http.StatusBadRequest, nil, errors.New("invalid shift data")
	}

	// Perform any validation or preprocessing of the shift data if needed

	// Save the user to the database
	if err := ss.db.Create(&user).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to create shift")
	}
	// Return the created user as the response
	return http.StatusOK, user, nil
}
