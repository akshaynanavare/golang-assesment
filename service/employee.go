package service

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/employee-management/model"
	repository "github.com/employee-management/repository/init"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func GetEmployeeByID(c *gin.Context) {
	log.Printf("recieved get employee request: %v", c.Request)

	request := c.Request.WithContext(c)

	employeeID := c.Query("employee_id")

	emp, err := repository.Postgresql.Employee.GetByID(request.Context(), employeeID)
	if err != nil {
		log.Printf("error getting employee from db, error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, emp)
	log.Print("successfully sent employee data")
}

func CreateEmployee() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Printf("recieved create employee request: %v", c.Request)

		request := c.Request.WithContext(c)

		var emp model.Employee

		// Bind JSON request body to the Employee struct
		if err := c.ShouldBindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uniqueID, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		emp.ID = uniqueID.String()

		err = repository.Postgresql.Employee.Upsert(request.Context(), &emp)
		if err != nil {
			log.Printf("error creating employee in db, error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, emp)
		log.Print("successfully saved employee data")
	}
}

func UpdateEmployee() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Printf("recieved create employee request: %v", c.Request)

		request := c.Request.WithContext(c)

		var emp model.Employee

		// Bind JSON request body to the Employee struct
		if err := c.ShouldBindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := repository.Postgresql.Employee.GetByID(request.Context(), emp.ID)
		if err != nil {
			if err != sql.ErrNoRows {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id received"})
				return
			}

			log.Printf("error getting employee from db, error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = repository.Postgresql.Employee.Upsert(request.Context(), &emp)
		if err != nil {
			log.Printf("error updating employee in db, error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, emp)
		log.Print("successfully saved employee data")
	}
}

func DeleteEmployee() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Printf("recieved get employee request: %v", c.Request)

		request := c.Request.WithContext(c)

		employeeID := c.Query("employee_id")

		err := repository.Postgresql.Employee.Delete(request.Context(), employeeID)
		if err != nil {
			log.Printf("error getting employee from db, error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
		log.Print("successfully deleted employee data")
	}
}

func GetEmployeeList() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Printf("recieved get employee request: %v", c.Request)

		request := c.Request.WithContext(c)

		page := cast.ToInt(c.Query("page"))
		limit := cast.ToInt(c.Query("limit"))

		if page <= 0 || limit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "page and limit must be greater than 0",
			})
			return
		}

		offset := (page - 1) * limit

		emp, err := repository.Postgresql.Employee.GetList(request.Context(), limit, offset)
		if err != nil {
			log.Printf("error getting employee from db, error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, emp)
		log.Print("successfully sent employee data")
	}
}
