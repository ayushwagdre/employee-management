package models

import (
	"time"
)

type Employee struct {
	Name      string    `gorm:"column:name"`
	Position  string    `gorm:"column:position"`
	Salary    float64   `gorm:"column:salary"`
	Active    *bool     `gorm:"column:active"`
	Code      string    `gorm:"<-:false"`
	Password  string    `gorm:"column:password_digest"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName specifies the table name for the Employee model
func (Employee) TableName() string {
	return "employees"
}
