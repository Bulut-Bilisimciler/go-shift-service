package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile godoc
// @Summary upload file
// @Schemes
// @Description upload file
// @Tags Upload
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce */*
// @Success 200 {object} string "uploaded file name"
// @Failure 401 {object} handlers.RespondJson "Unauthorized"
// @Failure 403 {object} handlers.RespondJson "Forbidden"
// @Failure 500 {object} handlers.RespondJson "Internal Server Error"
// @Router /upload [post]
func (mss *ShiftService) HandleUploadFile(c *gin.Context) (int, interface{}, error) {
	// read file from request
	fileHeader, _ := c.FormFile("file")
	file, err := fileHeader.Open()
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("file not found")
	}
	defer file.Close()

	// Todo...
	return http.StatusOK, "file uploaded successfully", nil
}
