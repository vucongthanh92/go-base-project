package httpcommon

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse[T any] struct {
	Success bool            `json:"success"`
	Data    *T              `json:"data"`
	Errors  []ErrorResponse `json:"errors"`
}

type PagingSuccessResponse[T any] struct {
	Success bool            `json:"success"`
	Data    []T             `json:"data"`
	Errors  []ErrorResponse `json:"errors"`
	Meta    Meta            `json:"meta"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field"`
	Code    string `json:"code"`
}

func NewSuccessResponse[T any](data T) SuccessResponse[T] {
	return SuccessResponse[T]{
		Data:    &data,
		Success: true,
		Errors:  nil,
	}
}

func NewPartialSuccess[T any](success bool, data T, errors []ErrorResponse) SuccessResponse[T] {
	return SuccessResponse[T]{
		Data:    &data,
		Success: success,
		Errors:  errors,
	}
}

// NewPagingSuccessResponse take the capacity of array data as limit
// if you want to specific limit you can pass it as parameter
func NewPagingSuccessResponse[T any](data []T, total int, additionalData any, limit ...int) PagingSuccessResponse[T] {
	capacity := cap(data)
	if len(limit) > 0 && limit[0] > 0 {
		capacity = limit[0]
	}
	red := total / capacity
	if total%capacity > 0 {
		red = red + 1
	}

	// compare value red with current_page from request by limit[1]
	// limit[1] = pageIndex
	var isLastPage = false
	switch {
	case len(data) < capacity:
		isLastPage = true
	case len(data) == total:
		isLastPage = true
	case len(limit) >= 2 && limit[1] == red:
		isLastPage = true
	}

	meta := Meta{
		TotalCount:     int64(total),
		IsLastPage:     isLastPage,
		NumPage:        red,
		AdditionalData: additionalData,
	}

	return PagingSuccessResponse[T]{
		Data:    data,
		Success: true,
		Meta:    meta,
	}
}

func NewErrorResponse(message string, code string, field string) SuccessResponse[any] {
	var errors []ErrorResponse
	error := ErrorResponse{
		Message: message,
		Field:   field,
		Code:    code,
	}
	errors = append(errors, error)
	return SuccessResponse[any]{
		Success: false,
		Data:    nil,
		Errors:  errors,
	}
}

func NewError(message string, code string, field string) []ErrorResponse {
	var errors []ErrorResponse
	error := ErrorResponse{
		Message: message,
		Field:   field,
		Code:    code,
	}
	errors = append(errors, error)
	return errors
}

func AddError(errOrigin []ErrorResponse, message string, code string, field string) []ErrorResponse {
	var errors []ErrorResponse
	err := ErrorResponse{
		Message: message,
		Field:   field,
		Code:    code,
	}
	errors = append(errOrigin, err)
	return errors
}

type ErrorDTO struct {
	IsSystemError bool   `json:"isSystemError"`
	Field         string `json:"field"`
	Error         error  `json:"error"`
}

func NewErrorDTO(isSystemError bool, field string, err error) ErrorDTO {
	return ErrorDTO{
		IsSystemError: isSystemError,
		Field:         field,
		Error:         err,
	}
}

func IsErrorWithoutNoRows(err error) bool {
	return err != nil && err != sql.ErrNoRows
}

func ExposeError(c *gin.Context, errorCommon ErrorDTO) {
	var (
		httpStatus int
		errCode    string
		errMsg     = errorCommon.Error.Error()
	)

	switch {
	case !errorCommon.IsSystemError:
		{
			httpStatus = http.StatusBadRequest
			errCode = RequestInvalid
		}
	default:
		httpStatus = http.StatusInternalServerError
		errCode = SystemError
	}

	c.JSON(httpStatus, NewErrorResponse(errMsg, errCode, errorCommon.Field))
}
