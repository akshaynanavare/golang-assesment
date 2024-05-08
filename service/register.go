package service

import (
	"github.com/gin-gonic/gin"
)

func RegisterAPIs(router *gin.Engine) {
	router.GET("/employee", GetEmployeeByID())
	router.POST("/employee", CreateEmployee())
	router.PUT("/employee", UpdateEmployee())
	router.DELETE("/employee", DeleteEmployee())
	router.GET("/employees", GetEmployeeList())
}
