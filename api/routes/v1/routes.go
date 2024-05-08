package v1

import "github.com/julienschmidt/httprouter"

func Init(router *httprouter.Router) {
	InitMerchantRoutes(router)
}
