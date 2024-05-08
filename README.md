# golang-assesment

## Run migration script on the postgresql server
migrate -source file://migration -database "postgres://<username>:<password>@<hostname>:5432/postgres?sslmode=disable" up

## Steps to follow run follow commands
1. go mod tidy
2. go run cmd/main.go


## hierarchy of directories
1. cmd/main.go
    - This is driving command function
    - It initializes the http server and registers all APIs
2. constants
    - It has all the required constants which is used by other packages
3. migration
    - It contains all the DB migrations scripts which is required for project
4. mocks
    - Its auto generated interface's mock functions to drive unit tests smoothly
5. model
    - It has all db/json level structs
6. postgresql
    - It has postgresql queries go files
7. repository
    - This service follows repository pattern in which we list all table queries interfaces
    - where postgresql package implements all those queries and binds this interface to the database
8. service
    - This has all registered API handlers along with unit testcases
    - Due to time constraints, I have not covered error UT cases
