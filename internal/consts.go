package internal

import (
	"errors"
)

var ErrFlagNotFound = errors.New("flag not found")
var ErrFlagNotType = errors.New("flag is not of the expected type")
