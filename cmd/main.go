package main

import (
	"context"

	"github.com/employee-management/constants"
	repository "github.com/employee-management/repository/init"
	"github.com/employee-management/service"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	// open db connection
	repository.Connect(ctx, constants.Postgres)

	// Create a new Gin router
	router := gin.Default()

	service.RegisterAPIs(router)

	// Start the Gin server on port 8080
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
