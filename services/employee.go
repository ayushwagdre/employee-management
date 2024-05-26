package services

//go:generate mockgen -source=./merchant.go -destination=./mock_services/mock_merchant.go -package=mock_services

import (
	"context"
	"fmt"
	"sync"

	"practice/lib/errors"
	"practice/models"
	repository "practice/repositories"
	"practice/utils"
)

type GetEmployeeDetails struct {
	Name     string
	Position string
	Salary   float64
	Active   bool
	Code     string
}

type UpsertEmployeeDetailsOpts struct {
	Name     string
	Position string
	Salary   float64
	Active   *bool
	Password string
}

type EmployeeService interface {
	Create(context.Context, *UpsertEmployeeDetailsOpts) error
	Get(context.Context, string) (GetEmployeeDetails, error)
	GetAll(context.Context, int, int) ([]GetEmployeeDetails, error)
	Update(context.Context, *UpsertEmployeeDetailsOpts, string) error
}

type employeeService struct {
	employeeRepo repository.EmployeeRepository
	mutex        sync.Mutex // Mutex for synchronization
}

func NewEmployeeService() EmployeeService {
	return &employeeService{
		employeeRepo: repository.NewEmployeeRepository(),
	}
}

func (m *employeeService) Create(ctx context.Context, opts *UpsertEmployeeDetailsOpts) error {

	hashPassword, err := utils.HashPassword(opts.Password)
	if err != nil {
		fmt.Println(err)
	}

	params := &models.Employee{
		Name:     opts.Name,
		Position: opts.Position,
		Salary:   opts.Salary,
		Active:   opts.Active,
		Password: hashPassword,
	}

	err = m.employeeRepo.Create(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (m *employeeService) Get(ctx context.Context, code string) (GetEmployeeDetails, error) {
	employeeDetail, err := m.employeeRepo.Get(ctx, code)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return GetEmployeeDetails{}, errors.ErrRecordNotFound
		}
		return GetEmployeeDetails{}, err
	}

	return GetEmployeeDetails{
		Name:     employeeDetail.Name,
		Position: employeeDetail.Position,
		Salary:   employeeDetail.Salary,
		Active:   *employeeDetail.Active,
		Code:     employeeDetail.Code,
	}, nil
}

func (m *employeeService) GetAll(ctx context.Context, offset, limit int) ([]GetEmployeeDetails, error) {
	employeeDetail, err := m.employeeRepo.GetAll(ctx, offset, limit)
	if err != nil {
		return []GetEmployeeDetails{}, err
	}

	var employeeDetails []GetEmployeeDetails
	for _, employee := range employeeDetail {
		employeeDetails = append(employeeDetails, GetEmployeeDetails{
			Name:     employee.Name,
			Position: employee.Position,
			Salary:   employee.Salary,
			Active:   *employee.Active,
			Code:     employee.Code,
		})
	}
	return employeeDetails, nil
}

func (m *employeeService) Update(ctx context.Context, updateOpts *UpsertEmployeeDetailsOpts, employeeCode string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	params := &models.Employee{
		Code: employeeCode,
	}

	if params.Name != "" {
		params.Name = updateOpts.Name
	}
	if params.Position != "" {
		params.Position = updateOpts.Position
	}
	if params.Salary != 0 {
		params.Salary = updateOpts.Salary
	}
	if params.Active != nil {
		params.Active = updateOpts.Active
	}
	if params.Password != "" {
		hashPassword, err := utils.HashPassword(params.Password)
		if err != nil {
			return err
		}
		params.Password = hashPassword
	}

	err := m.employeeRepo.Update(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
