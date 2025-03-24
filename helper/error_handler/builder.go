package errorhandler

import (
	"github.com/gin-gonic/gin"
	httpcommon "github.com/vucongthanh92/go-base-project/helper/http_common"
	"github.com/vucongthanh92/go-base-project/helper/utils"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
)

type ErrorBuilder struct {
	IsSystemError   bool              `json:"is_system_error"`
	IsMultipleError bool              `json:"is_multiple_error"`
	LogError        error             `json:"log_error"`
	Status          int               `json:"status"`
	Errors          []models.ErrorDTO `json:"errors"`
}

func InitErrorBuilder() *ErrorBuilder {
	return &ErrorBuilder{}
}

// setup error ----------------------------------------------------------

func (b *ErrorBuilder) SetIsSystemError(req bool) *ErrorBuilder {
	b.IsSystemError = req
	return b
}

func (b *ErrorBuilder) SetIsMultipleError(req bool) *ErrorBuilder {
	b.IsMultipleError = req
	return b
}

func (b *ErrorBuilder) SetLogError(req error) *ErrorBuilder {
	b.LogError = req
	return b
}

func (b *ErrorBuilder) SetError(req models.ErrorDTO) *ErrorBuilder {
	b.Errors = append(b.Errors, req)
	return b
}

func (b *ErrorBuilder) SetStatus(req int) *ErrorBuilder {
	b.Status = req
	return b
}

// expose error ----------------------------------------------------------

func (b *ErrorBuilder) ExposeHttpError(c *gin.Context) {

	errors := []models.ErrorDTO{}

	utils.IterateSlice(b.Errors, func(i int, err models.ErrorDTO) {
		errors = append(errors, err)
	})

	response := httpcommon.SuccessResponse[any]{
		Success: false,
		Data:    nil,
		Errors:  errors,
	}

	c.JSON(b.Status, response)
}
