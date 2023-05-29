package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	_ "github.com/lib/pq"
=======
	"gorm.io/gorm"
>>>>>>> babb95df0fea093f40084966abceab22a79853b9
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
<<<<<<< HEAD
	// var params models.Pagination
	// if err := c.ShouldBindQuery(&params); !errors.Is(err, nil) {
	// 	return http.StatusBadRequest, nil, errors.New("invalid page or limit query for pagination")
	// }

	// limit := params.Limit
	// offset := (params.Page - 1) * params.Limit

	// get users
	// var users []models.User
	// if err := ss.db.Where("deleted_at is NULL").Order("created_at DESC").Find(&users).Error; !errors.Is(err, nil) {
	// 	return http.StatusNotFound, nil, errors.New("users not found")
	// }
	db, err := sql.Open("postgres", "postgres://shiftuser:shiftdb@localhost:5432/shiftdb?sslmode=disable")
	if err != nil {
		log.Print("err")
=======

	// get users
	var users []models.User
	err := ss.db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusOK, users, errors.New("no users found")
		}

>>>>>>> babb95df0fea093f40084966abceab22a79853b9
		return http.StatusInternalServerError, nil, err
	}
	defer db.Close()

<<<<<<< HEAD
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, nil, err
	}
	defer rows.Close()
	var data []models.User
=======
	return http.StatusOK, users, nil
>>>>>>> babb95df0fea093f40084966abceab22a79853b9

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.Made_Field,
			&user.Nickname,
			&user.UserId,
			&user.Creation_Date,
			&user.Update_Date,
		)
		if err != nil {
			log.Print(err)
			return http.StatusInternalServerError, nil, err
		}
		data = append(data, user) // Kullanıcıyı slice'a ekle
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, data, nil
}
