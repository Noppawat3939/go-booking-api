package domain

import "errors"

var (
	ErrInternal = errors.New("internal server error")
	ErrConflict = errors.New("data conflicts with existing data in unique column")
)
