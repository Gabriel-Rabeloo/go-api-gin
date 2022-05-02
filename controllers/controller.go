package controllers

import (
	"github.com/Gabriel-Rabeloo/go-api-gin/models"
	"github.com/gin-gonic/gin"
)

func FillAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Abc(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API diz:": "asbdadsa" + name + "ashdajdas",
	})
}
