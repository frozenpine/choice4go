package choice4go

import "errors"

var (
	ErrUnsupportedSys     = errors.New("unsupported system")
	ErrLoadLib            = errors.New("load library failed")
	ErrLoadFunc           = errors.New("load function failed")
	ErrInitialized        = errors.New("choice api not initialized")
	ErrMainCbFailed       = errors.New("set main callback failed")
	ErrDataEmpty          = errors.New("data is empty")
	ErrDataLenMissMatch   = errors.New("data len mismatch")
	ErrStart              = errors.New("choice start failed")
	ErrStop               = errors.New("choice stop failed")
	ErrGetData            = errors.New("make data structure failed")
	ErrEQCall             = errors.New("choice func call failed")
	ErrInvalidArgs        = errors.New("choice func call with invalid args")
	ErrInvalidOptionValue = errors.New("unsupport option value type")

	ErrInvalidReportName  = errors.New("invalid report name value")
	ErrInvalidReportType  = errors.New("invalid report type value")
	ErrInvalidCompanyType = errors.New("invalid company type value")
	ErrInvalidSourceType  = errors.New("invalid report source value")
	ErrInvalidDataValue   = errors.New("invalid data value")
)
