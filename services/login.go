package services

//go:generate mockgen -source=./merchant.go -destination=./mock_services/mock_merchant.go -package=mock_services

import (
	"context"
	"sync"

	"practice/lib/errors"
	repository "practice/repositories"
	"practice/utils"
)

type LoginDetailsOpts struct {
	Code     string
	Password string
}

type LoginService interface {
	Login(context.Context, *LoginDetailsOpts) error
}

type loginService struct {
	employeeRepo repository.EmployeeRepository
	mutex        sync.Mutex // Mutex for synchronization
}

func NewLoginService() LoginService {
	return &loginService{
		employeeRepo: repository.NewEmployeeRepository(),
	}
}

func (m *loginService) Login(ctx context.Context, opts *LoginDetailsOpts) error {
	employeeDetails, err := m.employeeRepo.Get(ctx, opts.Code)
	if err != nil {
		if errors.Is(repository.ErrRecordNotFound, err) {
			return errors.Wrapf(errors.ErrRecordNotFound, "employee with code %s not found", opts.Code)
		}
	}

	if utils.CheckPasswordHash(opts.Password, employeeDetails.Password) {
		return nil
	}
	return errors.Wrapf(errors.ErrBadRequest, "invalid password")
}
