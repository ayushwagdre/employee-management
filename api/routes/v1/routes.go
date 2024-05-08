package v1

import (
	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
)

func Init(router *nrhttprouter.Router) {
	InitMerchantRoutes(router)
}
