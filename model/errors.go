package model

import "errors"

var ErrNotFound = errors.New("not found")

const (
	IDNotANumberMessageErr = "%s is not a number"
	UnexpectedErrOccurred  = "Unexpected error occurred: %s"
	DataNotFoundErr        = "%s not found"
	BadRequestErr          = "%s is invalid"
)
