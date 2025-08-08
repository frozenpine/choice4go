package choice4go

import "errors"

var (
	ErrUnsupportedSys   = errors.New("unsupported system")
	ErrLoadLib          = errors.New("load library failed")
	ErrLoadFunc         = errors.New("load function failed")
	ErrInitialized      = errors.New("choice api not initialized")
	ErrMainCbFailed     = errors.New("set main callback failed")
	ErrDataEmpty        = errors.New("data is empty")
	ErrDataLenMissMatch = errors.New("data len mismatch")
	ErrStart            = errors.New("choice start failed")
	ErrStop             = errors.New("choice stop failed")
	ErrGetData          = errors.New("make data structure failed")
	ErrEQCall           = errors.New("choice func call failed")
	ErrInvalidArgs      = errors.New("choice func call with invalid args")
)
