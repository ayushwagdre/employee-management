package endpoints

import (
	"github.com/jinzhu/copier"

	"practice/lib/errors"
	"practice/lib/web"
	"practice/services"
)

type upsertDetailsRequest struct {
	Name     string  `json:"name" validate:"required"`
	Position string  `json:"position" validate:"required"`
	Salary   float64 `json:"salary" validate:"required"`
	Active   *bool   `json:"active"`
}

type getOrUpdateDetailsRequest struct {
	Code string `json:"code" validate:"required"`
}

type getAllDetailsRequest struct {
	Offset int `json:"offset,string" validate:"required"`
	Limit  int `json:"limit,string" validate:"required"`
}

type upsertDetailsResponse struct {
	Message string `json:"message"`
}

type getDetailsResponse struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	Active   bool    `json:"active"`
	Code     string  `json:"code"`
}

type EmployeeEndpoints interface {
	Create(*web.Request) web.Response
	Get(*web.Request) web.Response
	GetAll(*web.Request) web.Response
	Update(*web.Request) web.Response
}

type employeeEndpoints struct {
	service services.EmployeeService
}

func NewEmployeeEndpoints() EmployeeEndpoints {
	return &employeeEndpoints{
		service: services.NewEmployeeService(),
	}
}

func (e *employeeEndpoints) Create(req *web.Request) web.Response {
	params := &upsertDetailsRequest{}
	err := req.ParseAndValidateBody(params)
	if err != nil {
		return e.errorResponse(err)
	}

	employeeDetails := &services.UpsertEmployeeDetailsOpts{}
	copier.Copy(employeeDetails, params)

	err = e.service.Create(req.Context(), employeeDetails)
	if err != nil {
		return e.errorResponse(err)
	}

	resp := &upsertDetailsResponse{
		Message: "upserted successfully",
	}
	return web.NewSuccessResponse(resp, web.StatusOK, web.V1Api)
}

func (e *employeeEndpoints) Get(req *web.Request) web.Response {
	var (
		params   getOrUpdateDetailsRequest
		employee services.GetEmployeeDetails
		err      error
	)

	if err = req.ParseAndValidateParams(&params); err != nil {
		return e.errorResponse(err)
	}

	employee, err = e.service.Get(req.Context(), params.Code)

	if err != nil {
		return e.errorResponse(err)
	}

	resp := &getDetailsResponse{}
	copier.Copy(resp, employee)

	return web.NewSuccessResponse(resp, web.StatusOK, web.V1Api)
}

func (e *employeeEndpoints) GetAll(req *web.Request) web.Response {
	var (
		params   getAllDetailsRequest
		employee []services.GetEmployeeDetails
		err      error
	)

	if err = req.ParseAndValidateParams(&params); err != nil {
		return e.errorResponse(err)
	}

	employee, err = e.service.GetAll(req.Context(), params.Offset, params.Limit)

	if err != nil {
		return e.errorResponse(err)
	}

	resp := []getDetailsResponse{}
	copier.Copy(&resp, employee)

	return web.NewSuccessResponse(resp, web.StatusOK, web.V1Api)
}

func (e *employeeEndpoints) Update(req *web.Request) web.Response {
	params := &upsertDetailsRequest{}
	err := req.ParseAndValidateBody(params)
	if err != nil {
		return e.errorResponse(err)
	}

	employeeCode := req.GetPathParam("code")
	if employeeCode == "" {
		return web.ErrBadRequest("code is required", web.V1Api)
	}

	employeeDetails := &services.UpsertEmployeeDetailsOpts{}
	copier.Copy(employeeDetails, params)

	err = e.service.Update(req.Context(), employeeDetails, employeeCode)

	if err != nil {
		return e.errorResponse(err)
	}

	resp := &upsertDetailsResponse{
		Message: "upserted successfully",
	}
	return web.NewSuccessResponse(resp, web.StatusOK, web.V1Api)
}

func (e *employeeEndpoints) errorResponse(err error) web.Response {
	switch {
	case errors.Is(err, errors.ErrBadRequest):
		return web.ErrBadRequest(err.Error(), web.V1Api)
	case errors.Is(err, ErrRecordNotFound):
		return web.NewError(errors.Original(err).Code(), "merchant partner employee not found", web.StatusBadRequest, web.V1Api)
	default:
		return web.ErrInternalServerError(err.Error(), web.V1Api)
	}
}