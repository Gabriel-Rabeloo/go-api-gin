package controllers

import (
	"fmt"
	"net/http"

	"github.com/Gabriel-Rabeloo/go-api-gin/database"
	"github.com/Gabriel-Rabeloo/go-api-gin/models"
	"github.com/gin-gonic/gin"
)

func FindAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func FindByIdStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})

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

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	fmt.Println()
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error()})
		return
	}

	if err := database.DB.Model(&student).Where("id = ?", id).UpdateColumns(student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)

	c.JSON(http.StatusNoContent, gin.H{})
}
