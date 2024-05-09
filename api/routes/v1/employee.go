package v1

import (
	"practice/api/endpoints"
	"practice/api/middlewares"
	"practice/lib/web"

	"github.com/julienschmidt/httprouter"
)

const (
	createEmployeeEndpoint = "/api/v1/employee"
	getAllEmployeeEndpoint = "/api/v1/employee"
	getEmployeeEndpoint    = "/api/v1/employee/:code"
	updateEmployeeEndpoint = "/api/v1/employee/:code"
)

func InitEmployeeRoutes(router *httprouter.Router) {
	employeeEndpoints := endpoints.NewEmployeeEndpoints()
	router.POST(createEmployeeEndpoint, web.Serve(employeeEndpoints.Create, middlewares.DefaultMiddlewares...))
	router.GET(getEmployeeEndpoint, web.Serve(employeeEndpoints.Get, middlewares.DefaultMiddlewares...))
	router.GET(getAllEmployeeEndpoint, web.Serve(employeeEndpoints.GetAll, middlewares.DefaultMiddlewares...))
	router.PUT(updateEmployeeEndpoint, web.Serve(employeeEndpoints.Update, middlewares.DefaultMiddlewares...))
}
