package errorCodes

//go:generate stringer -type=ErrorCode
type ErrorCode uint

const (
	Undefined ErrorCode = iota
	Forbidden           //403 forbidden
	NotFound            //404 not found
	BadRequest
	DataValidationError // 400 bad request
	ConflictingData     //409 conflict
	InternalError       //500 internal server error
)
