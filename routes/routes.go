package routes

import (
	"github.com/Gabriel-Rabeloo/go-api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/students", controllers.FindAllStudents)
	r.POST("/api/students", controllers.CreateStudent)
	r.GET("/api/students/:id", controllers.FindByIdStudent)
	r.PATCH("/api/students/:id", controllers.UpdateStudent)
	r.DELETE("/api/students/:id", controllers.DeleteStudent)
	r.Run()

}
