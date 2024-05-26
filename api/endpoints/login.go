package endpoints

import (
	"github.com/jinzhu/copier"

	"practice/lib/errors"
	"practice/lib/web"
	"practice/services"
)

type loginDetailsRequest struct {
	Code     string `json:"code" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginDetailsResponse struct {
	Message string `json:"message"`
}

type LoginEndpoints interface {
	Login(*web.Request) web.Response
}

type loginEndpoints struct {
	service services.LoginService
}

func NewLoginEndpoints() LoginEndpoints {
	return &loginEndpoints{
		service: services.NewLoginService(),
	}
}

func (e *loginEndpoints) Login(req *web.Request) web.Response {
	params := &loginDetailsRequest{}
	err := req.ParseAndValidateBody(params)
	if err != nil {
		return e.errorResponse(err)
	}

	loginDetails := &services.LoginDetailsOpts{}
	copier.Copy(loginDetails, params)

	err = e.service.Login(req.Context(), loginDetails)
	if err != nil {
		return e.errorResponse(err)
	}

	resp := &loginDetailsResponse{
		Message: "login successfully",
	}
	return web.NewSuccessResponse(resp, web.StatusOK, web.V1Api)
}

func (e *loginEndpoints) errorResponse(err error) web.Response {
	switch {
	case errors.Is(err, errors.ErrBadRequest):
		return web.ErrBadRequest(err.Error(), web.V1Api)
	case errors.Is(err, errors.ErrRecordNotFound):
		return web.NewError(errors.Original(err).Code(), "employee record not found", web.StatusBadRequest, web.V1Api)
	default:
		return web.ErrInternalServerError(err.Error(), web.V1Api)
	}
}
