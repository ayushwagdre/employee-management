package endpoints

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"practice/lib/errors"
	"practice/lib/web"
	"practice/services"
)

type upsertPartnerAppsConfigRequest struct {
	MerchantID    uuid.UUID   `json:"merchant_id" validate:"uuid,required"`
	ClientID      uuid.UUID   `json:"client_id" validate:"uuid,required"`
	PartnerConfig interface{} `json:"partner_apps_config" validate:"required"`
}

type getPartnerAppsConfigRequest struct {
	MerchantID *uuid.UUID `json:"merchant_id" validate:"omitempty,uuid"`
	ClientID   *uuid.UUID `json:"client_id" validate:"omitempty,uuid"`
}

type upsertPartnerAppsConfigResponse struct {
	Message string `json:"message"`
}

type getPartnerAppsConfigResponse struct {
	PartnerAppsConfig interface{} `json:"partner_apps_config"`
}

type MerchantEndpoints interface {
	UpsertPartnerAppsConfig(*web.Request) web.Response
	GetPartnerAppsConfig(*web.Request) web.Response
}

type merchantEndpoints struct {
	service services.MerchantService
}

func NewMerchantEndpoints() MerchantEndpoints {
	return &merchantEndpoints{
		service: services.NewMerchantService(),
	}
}

func (e *merchantEndpoints) UpsertPartnerAppsConfig(req *web.Request) web.Response {
	params := &upsertPartnerAppsConfigRequest{}
	err := req.ParseAndValidateBody(params)
	if err != nil {
		return e.errorResponse(err)
	}

	partnerAppsConfig := services.PartnerAppsConfig{}
	copier.Copy(&partnerAppsConfig, params)

	err = e.service.UpsertPartnerAppsConfig(req.Context(), partnerAppsConfig)
	if err != nil {
		return e.errorResponse(err)
	}

	resp := &upsertPartnerAppsConfigResponse{
		Message: "upserted successfully",
	}
	return web.NewSuccessResponse(resp, web.StatusOK, web.V1Api)
}

func (e *merchantEndpoints) GetPartnerAppsConfig(req *web.Request) web.Response {
	var (
		params getPartnerAppsConfigRequest
		config services.MerchantPartnerAppsConfig
		err    error
	)

	if err = req.ParseAndValidateParams(&params); err != nil {
		return e.errorResponse(err)
	}
	if params.MerchantID == nil && params.ClientID == nil {
		return e.errorResponse(errors.Wrapf(errors.ErrBadRequest, "merchant_id or client_id is required"))
	}

	config, err = e.service.GetPartnerAppsConfigByMerchantID(req.Context(), *params.MerchantID)

	if err != nil {
		return e.errorResponse(err)
	}

	resp := &getPartnerAppsConfigResponse{}
	copier.Copy(resp, config)

	return web.NewSuccessResponse(resp, web.StatusOK, web.V1Api)
}

func (e *merchantEndpoints) errorResponse(err error) web.Response {
	switch {
	case errors.Is(err, errors.ErrBadRequest):
		return web.ErrBadRequest(err.Error(), web.V1Api)
	case errors.Is(err, ErrRecordNotFound):
		return web.NewError(errors.Original(err).Code(), "merchant partner config not found", web.StatusBadRequest, web.V1Api)
	default:
		return web.ErrInternalServerError(err.Error(), web.V1Api)
	}
}
