package endpoints

import (
	"practice/lib/errors"
)

var (
	ErrInvalidIntegration = errors.NewWithCode("invalid_integration")
	ErrBadRequest         = errors.NewWithCode("bad_request")
	ErrRecordNotFound     = errors.NewWithCode("record_not_found")
)
