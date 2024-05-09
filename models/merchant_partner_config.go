package models

import (
	"time"
)

type Employee struct {
	Name      string
	Position  string
	Salary    float64
	Active    *bool
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
