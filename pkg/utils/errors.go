package utils

import "errors"

var (
	ErrEmptyAlphabet = errors.New("empty alphabet")
	ErrInvalidLength = errors.New("invalid length")
)
