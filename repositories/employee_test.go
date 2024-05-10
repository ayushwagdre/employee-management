package repository

import (
	"context"
	"practice/lib/db"
	"practice/models"
	"testing"

	"github.com/stretchr/testify/assert"

	"practice/test_helpers"
)

func TestCheckoutConfigRepository_Upsert(t *testing.T) {
	repo := NewEmployeeRepository()
	active := true
	tests := []struct {
		name   string
		params *models.Employee
	}{
		{
			name: "should successfully create or update the ui_configs",
			params: &models.Employee{
				Name:     "test",
				Position: "test",
				Salary:   1000,
				Active:   &active,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Create(context.Background(), tc.params)
			assert.NoError(t, err, "Create should not return an error")
		})
		test_helpers.ClearDataFromPostgres(db.Get().Get())
	}
}

func TestCheckoutConfigRepository_Get(t *testing.T) {
	repo := NewEmployeeRepository()
	active := true
	tests := []struct {
		name         string
		employeeCode string
		create       *models.Employee
	}{
		{
			name:         "should successfully fetch all the employee by code",
			employeeCode: "EMP1001",
			create: &models.Employee{
				Name:     "test",
				Position: "test",
				Salary:   1000,
				Active:   &active,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.Create(context.Background(), tc.create)
			employee, err := repo.Get(context.Background(), tc.employeeCode)
			assert.NoError(t, err, "get employee should not return an error")
			assert.Len(t, employee, 1)
			test_helpers.ClearDataFromPostgres(db.Get().Get())
		})
	}
}

func TestCheckoutConfigRepository_GetAll(t *testing.T) {
	repo := NewEmployeeRepository()
	active := true
	tests := []struct {
		name         string
		employeeCode string
		create       *models.Employee
	}{
		{
			name:         "should successfully fetch all the employee by code",
			employeeCode: "EMP1001",
			create: &models.Employee{
				Name:     "test",
				Position: "test",
				Salary:   1000,
				Active:   &active,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.Create(context.Background(), tc.create)
			repo.Create(context.Background(), tc.create)
			repo.Create(context.Background(), tc.create)
			employee, err := repo.GetAll(context.Background(), 0, 10)
			assert.NoError(t, err, "get employee should not return an error")
			assert.Len(t, employee, 3)
			test_helpers.ClearDataFromPostgres(db.Get().Get())
		})
	}
}

func TestCheckoutConfigRepository_Update(t *testing.T) {
	repo := NewEmployeeRepository()
	active := true
	code := "EMP1001"
	tests := []struct {
		name   string
		params *models.Employee
	}{
		{
			name: "should successfully create or update the ui_configs",
			params: &models.Employee{
				Name:     "test",
				Position: "test",
				Salary:   1000,
				Active:   &active,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.Create(context.Background(), tc.params)
			tc.params.Code = code
			tc.name = tc.name + "update"
			err := repo.Update(context.Background(), tc.params)
			assert.NoError(t, err, "Create should not return an error")
		})
		test_helpers.ClearDataFromPostgres(db.Get().Get())
	}
}
