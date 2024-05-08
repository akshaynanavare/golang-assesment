package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/employee-management/mocks/repository/intf"
	"github.com/employee-management/model"
	repository "github.com/employee-management/repository/init"
	"github.com/employee-management/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var router *gin.Engine

func testSuite(t *testing.T) func(t *testing.T) {
	repository.SetPostgresMock()
	router = gin.Default()
	service.RegisterAPIs(router)
	// Return a function to teardown the test
	return func(t *testing.T) {
		repository.SetPostgresMock()
	}
}

func TestEmployeeAllAPIs(t *testing.T) {
	defer testSuite(t)(t)
	type args struct {
		c      *gin.Context
		method string
		path   string
		body   map[string]interface{}
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantResponse   map[string]interface{}
		mockRepo       func()
	}{
		{
			name: "should success get employee API when valid id is passed",
			args: args{
				c: &gin.Context{
					Params: gin.Params{
						gin.Param{
							Key:   "employee_id",
							Value: "123",
						},
					},
				},
				method: "GET",
				path:   "/employee?employee_id=123",
				body:   make(map[string]interface{}),
			},
			wantStatusCode: http.StatusOK,
			wantResponse: map[string]interface{}{
				"id":       "123",
				"name":     "John Doe",
				"position": "Software Engineer",
				"salary":   float64(100000),
			},
			mockRepo: func() {
				emp := mocks.NewEmployee(t)
				emp.EXPECT().GetByID(mock.Anything, "123").Return(&model.Employee{
					ID:       "123",
					Name:     "John Doe",
					Position: "Software Engineer",
					Salary:   100000,
				}, nil)
				repository.Postgresql.Employee = emp
			},
		},
		{
			name: "should success PUT employee API when valid id is passed for updates",
			args: args{
				c:      &gin.Context{},
				method: "PUT",
				path:   "/employee",
				body: map[string]interface{}{
					"id":       "123",
					"name":     "John Doe",
					"position": "Software Engineer",
					"salary":   100000,
				},
			},
			wantStatusCode: http.StatusOK,
			wantResponse: map[string]interface{}{
				"id":       "123",
				"name":     "John Doe",
				"position": "Software Engineer",
				"salary":   float64(100000),
			},
			mockRepo: func() {
				emp := mocks.NewEmployee(t)
				emp.EXPECT().GetByID(mock.Anything, "123").Return(&model.Employee{
					ID:       "123",
					Name:     "John Doe",
					Position: "Software Engineer",
					Salary:   100000,
				}, nil)
				emp.EXPECT().Upsert(mock.Anything, mock.Anything).Return(nil)
				repository.Postgresql.Employee = emp
			},
		},
		{
			name: "should delete an employee success when valid id is passed",
			args: args{
				c: &gin.Context{
					Params: gin.Params{
						gin.Param{
							Key:   "employee_id",
							Value: "123",
						},
					},
				},
				method: "DELETE",
				path:   "/employee",
				body:   make(map[string]interface{}),
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   nil,
			mockRepo: func() {
				emp := mocks.NewEmployee(t)
				emp.EXPECT().Delete(mock.Anything, mock.Anything).Return(nil)
				repository.Postgresql.Employee = emp
			},
		},
		{
			name: "should create an employee when valid request is passed",
			args: args{
				c:      &gin.Context{},
				method: "POST",
				path:   "/employee",
				body: map[string]interface{}{
					"name":     "John Doe",
					"position": "Software Engineer",
					"salary":   100000,
				},
			},
			wantStatusCode: http.StatusCreated,
			wantResponse: map[string]interface{}{
				"name":     "John Doe",
				"position": "Software Engineer",
				"salary":   float64(100000),
			},
			mockRepo: func() {
				emp := mocks.NewEmployee(t)
				emp.EXPECT().Upsert(mock.Anything, mock.Anything).Return(nil)
				repository.Postgresql.Employee = emp
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockRepo()
			// Create a mock request
			body, err := json.Marshal(tt.args.body)
			if err != nil {
				t.Fatalf("failed to marshal request body: %v", err)
			}

			req, err := http.NewRequest(tt.args.method, tt.args.path, bytes.NewBuffer([]byte(body)))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			// Check the status code
			assert.Equal(t, tt.wantStatusCode, rr.Code)

			if tt.wantResponse != nil {
				// Parse the response body
				var response map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("failed to parse response body: %v", err)
				}

				// assert response body
				assert.Equal(t, tt.wantResponse["name"], response["name"])
				assert.Equal(t, tt.wantResponse["salary"], response["salary"])
				assert.Equal(t, tt.wantResponse["position"], response["position"])
			}
		})
	}
}
