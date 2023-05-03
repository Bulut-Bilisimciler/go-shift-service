package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// RetrieveFile godoc
// @Summary retrieve uploaded file by filename
// @Schemes
// @Description retrieve uploaded file by filename
// @Tags Upload
// @Accept */*
// @Param filename path string true "filename"
// @Produce */*
// @Success 200 {object} string "file content"
// @Failure 401 {object} handlers.RespondJson "Unauthorized"
// @Failure 403 {object} handlers.RespondJson "Forbidden"
// @Failure 500 {object} handlers.RespondJson "Internal Server Error"
// @Router /uploads/{filename} [get]
func (mss *ShiftService) HandleRetrieveFile(c *gin.Context) {
	// get :filename from url param
	filename := c.Param("filename")
	log.Println(filename)

	// Todo...
}
