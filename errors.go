package choice4go

import "errors"

var (
	ErrUnsupportedSys    = errors.New("unsupported system")
	ErrLoadLib           = errors.New("load library failed")
	ErrLoadFunc          = errors.New("load function failed")
	ErrInitialized       = errors.New("choice api not initialized")
	ErrMainCbFailed      = errors.New("set main callback failed")
	ErrDataEmpty         = errors.New("data is empty")
	ErrDataLenMissMatch  = errors.New("data len mismatch")
	ErrEQCall            = errors.New("choice func call failed")
	ErrInvalidCodes      = errors.New("invalid codes")
	ErrTooManyIndicators = errors.New("too many indicators")
)
