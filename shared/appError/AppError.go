package appError

import (
	"errors"

	"github.com/Nebuska/neblab/shared/appError/errorCodes"
	"gorm.io/gorm"
)

type AppError struct {
	ErrorCode errorCodes.ErrorCode
	Source    string
	Message   string
}

func New(code errorCodes.ErrorCode, source, message string) AppError {
	return AppError{
		ErrorCode: code,
		Source:    source,
		Message:   message,
	}
}

func FromGormError(err error) error {
	if err == nil {
		return nil
	}
	appError := AppError{
		Message: err.Error(),
		Source:  "GORM",
	}
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		appError.ErrorCode = errorCodes.NotFound
		return appError
	case errors.Is(err, gorm.ErrDuplicatedKey):
		appError.ErrorCode = errorCodes.ConflictingData
		return appError
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		appError.ErrorCode = errorCodes.ConflictingData
		return appError
	case errors.Is(err, gorm.ErrInvalidData):
		appError.ErrorCode = errorCodes.DataValidationError
		return appError
	}
	appError.ErrorCode = errorCodes.InternalError
	return appError
}

func (appError AppError) Error() string {
	return appError.Source + ": " + appError.Message
}
