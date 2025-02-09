package repository

import "errors"

var (
	ErrLinkNotFound      = errors.New("link with this code not found")
	ErrLinkAlreadyExists = errors.New("link related to this code or url already exists")
)
