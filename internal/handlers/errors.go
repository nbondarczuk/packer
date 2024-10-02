package handlers

import "errors"

var (
	ErrEmptyTagId = errors.New("an empty tag was provided")
)
