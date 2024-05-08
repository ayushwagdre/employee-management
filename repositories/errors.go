package repository

import "practice/lib/errors"

var (
	ErrRecordNotFound = errors.NewWithCode("record_not_found")
	ErrUnexpected     = errors.NewWithCode("unexpected")
)
