package controllers

import (
	"net/http"

	"github.com/Gabriel-Rabeloo/go-api-gin/database"
	"github.com/Gabriel-Rabeloo/go-api-gin/models"
	"github.com/gin-gonic/gin"
)

func FillAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func FindByIdStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.Data(404, gin.MIMEHTML, nil)

		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}
