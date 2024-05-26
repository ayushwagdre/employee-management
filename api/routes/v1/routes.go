package v1

import "github.com/julienschmidt/httprouter"

func Init(router *httprouter.Router) {
	InitEmployeeRoutes(router)
	InitEmployeeLoginRoutes(router)
}
