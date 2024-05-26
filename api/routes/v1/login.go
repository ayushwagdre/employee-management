package v1

import (
	"practice/api/endpoints"
	"practice/api/middlewares"
	"practice/lib/web"

	"github.com/julienschmidt/httprouter"
)

const (
	loginEmployeeEndpoint = "/api/v1/employee/login"
)

func InitEmployeeLoginRoutes(router *httprouter.Router) {
	loginEndpoints := endpoints.NewLoginEndpoints()
	router.POST(loginEmployeeEndpoint, web.Serve(loginEndpoints.Login, middlewares.DefaultMiddlewares...))
}
