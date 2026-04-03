package domain

import "errors"

var (
	ErrInvalidFromUserID = errors.New("invalid from user id")
	ErrInvalidToUserID   = errors.New("invalid to user id")
	ErrEmptyContent      = errors.New("content cannot be empty")
)
