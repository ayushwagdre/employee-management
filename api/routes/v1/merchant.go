package v1

import (
	"practice/api/endpoints"
	"practice/api/middlewares"
	"practice/lib/web"

	"github.com/julienschmidt/httprouter"
)

const (
	upsertPartnerAppsConfigEndpoint = "/internal/api/v1/merchant/partner_apps/config"
	getPartnerAppsConfigEndpoint    = "/internal/api/v1/merchant/partner_apps/config"
)

// InitMerchantPartnerRoutes initializes the routes related to merchant partner configuration.
func InitMerchantRoutes(router *httprouter.Router) {
	merchantEndpoints := endpoints.NewMerchantEndpoints()
	router.POST(upsertPartnerAppsConfigEndpoint, web.Serve(merchantEndpoints.UpsertPartnerAppsConfig, middlewares.DefaultMiddlewares...))
	router.GET(getPartnerAppsConfigEndpoint, web.Serve(merchantEndpoints.GetPartnerAppsConfig, middlewares.DefaultMiddlewares...))
}
