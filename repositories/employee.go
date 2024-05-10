package repository

import (
	"context"

	"practice/lib/db"
	"practice/lib/errors"
	"practice/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(ctx context.Context, params *models.Employee) error
	Get(ctx context.Context, code string) (*models.Employee, error)
	GetAll(ctx context.Context, offset, limit int) ([]*models.Employee, error)
	Update(ctx context.Context, params *models.Employee) error
}

type employeeRepository struct {
	db db.DB
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{db: db.Get()}
}

// duplicate employee can be handle phone number
func (r *employeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	err := r.db.Get().Create(&employee).Error
	if err != nil {
		return errors.Wrapf(ErrUnexpected, "failed to create employee: %s", err.Error())
	}
	return nil
}

func (r *employeeRepository) Get(ctx context.Context, code string) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.Get().Where("code = ?", code).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(ErrRecordNotFound, "employee with code %s not found", code)
		}
		return nil, errors.Wrapf(ErrUnexpected, "failed to get employee with code %s: %s", code, err.Error())
	}
	return &employee, nil
}

func (r *employeeRepository) GetAll(ctx context.Context, offset, limit int) ([]*models.Employee, error) {
	var employees []*models.Employee
	err := r.db.Get().Order("code").Offset(offset).Limit(limit).Find(&employees).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(ErrRecordNotFound, "no employee records found")
		}
		return nil, errors.Wrapf(ErrUnexpected, "failed to get employee records: %s", err.Error())
	}
	return employees, nil
}

func (r *employeeRepository) Update(ctx context.Context, updateOpts *models.Employee) error {
	err := r.db.Get().Model(&models.Employee{}).Where("code = ?", updateOpts.Code).Updates(updateOpts).Error
	if err != nil {
		return errors.Wrapf(ErrUnexpected, "failed to update employee with code %s: %s", updateOpts.Code, err.Error())
	}
	return nil
}
