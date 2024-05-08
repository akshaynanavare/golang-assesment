package service

import (
	"github.com/gin-gonic/gin"
)

func RegisterAPIs(router *gin.Engine) {
	router.GET("/employee", GetEmployeeByIDHandler())
	router.POST("/employee", CreateEmployee())
	router.PUT("/employee", UpdateEmployee())
	router.DELETE("/employee", DeleteEmployee())
	router.GET("/employees", GetEmployeeList())
}

func GetEmployeeByIDHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		GetEmployeeByID(c)
	}
}
