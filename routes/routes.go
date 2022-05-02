package routes

import (
	"github.com/Gabriel-Rabeloo/go-api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.FillAllStudents)
	r.GET("/:name", controllers.Abc)
	r.Run()

}
