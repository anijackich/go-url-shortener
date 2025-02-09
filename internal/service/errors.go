package service

import "errors"

var (
	ErrInvalidDomain = errors.New("invalid domain")
	ErrInvalidURL    = errors.New("invalid URL")
)
