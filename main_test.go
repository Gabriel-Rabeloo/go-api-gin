package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Gabriel-Rabeloo/go-api-gin/controllers"
	"github.com/Gabriel-Rabeloo/go-api-gin/database"
	"github.com/Gabriel-Rabeloo/go-api-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	database.Connect()

	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "test name", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGetStudents(t *testing.T) {
	r := SetupTestRoutes()
	CreateStudentMock()
	defer DeleteStudentMock()

	r.GET("/api/students", controllers.FindAllStudents)

	req, _ := http.NewRequest("GET", "/api/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Both status code should be equal")
}

func TestGetStudentById(t *testing.T) {
	r := SetupTestRoutes()
	CreateStudentMock()
	defer DeleteStudentMock()

	r.GET("/api/students/:id", controllers.FindByIdStudent)

	req, _ := http.NewRequest("GET", "/api/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var body models.Student
	json.Unmarshal(response.Body.Bytes(), &body)

	assert.Equal(t, http.StatusOK, response.Code, "Both status code should be equal")
	assert.Equal(t, body.Name, "test name", "Both names should be equal")
	assert.Equal(t, body.CPF, "12345678901", "Both cpf should be equal")
	assert.Equal(t, body.RG, "123456789", "Both rg should be equal")
}

func TestDeleteStudent(t *testing.T) {
	r := SetupTestRoutes()
	CreateStudentMock()

	r.DELETE("/api/students/:id", controllers.DeleteStudent)

	req, _ := http.NewRequest("DELETE", "/api/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNoContent, response.Code, "Both status code should be equal")
}

func TestUpdateStudent(t *testing.T) {
	r := SetupTestRoutes()
	CreateStudentMock()
	// defer DeleteStudentMock()

	r.PATCH("/api/students/:id", controllers.UpdateStudent)

	student := models.Student{Name: "test name - 2", CPF: "22222222222", RG: "123456789"}
	studentJson, _ := json.Marshal(student)

	req, _ := http.NewRequest("PATCH", "/api/students/"+strconv.Itoa(ID), bytes.NewBuffer(studentJson))
	fmt.Println(strconv.Itoa(ID))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var dbStudent models.Student
	database.DB.First(&dbStudent, ID)

	assert.Equal(t, http.StatusNoContent, response.Code, "Both status code should be equal")
	assert.Equal(t, student.Name, dbStudent.Name, "Both Name code should be equal")
	assert.Equal(t, student.CPF, dbStudent.CPF, "Both CPF code should be equal")
	assert.Equal(t, student.RG, dbStudent.RG, "Both RG code should be equal")
}
