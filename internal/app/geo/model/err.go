package model

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrOverflow = errors.New("overflow")
)
