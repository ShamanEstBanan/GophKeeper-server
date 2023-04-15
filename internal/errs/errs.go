package errs

import (
	"errors"
)

var (
	ErrInvalidAccessToken     = errors.New("invalid auth token")
	ErrInvalidLoginOrPassword = errors.New("login or password is invalid")
	ErrLoginIsEmpty           = errors.New("login is empty")
	ErrLoginAlreadyExist      = errors.New("login already exist")
	ErrPasswordIsEmpty        = errors.New("password is empty")
	ErrEmptyNameInRecord      = errors.New("empty name in record")
	ErrEmptyTypeInRecord      = errors.New("empty type in record")
	ErrEmptyDataInRecord      = errors.New("empty data in record")
	ErrNotFound               = errors.New("not found")
	ErrInvalidRecordID        = errors.New("invalid record ID")
	ErrInvalidRecordInfo      = errors.New("validation error")
)
