package validation

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/vucongthanh92/go-base-project/helper/constants"
	httpcommon "github.com/vucongthanh92/go-base-project/helper/http_common"
	"github.com/vucongthanh92/go-base-project/helper/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vucongthanh92/go-base-utils/tracing"
)

func CheckValidateHTTP(c *gin.Context, dest interface{}) (err error) {
	_, span := tracing.StartSpanFromContext(c.Request.Context(), "check_validate_http")
	defer span.End()
	if err = c.ShouldBindJSON(&dest); err != nil {
		checkErr(c, err)
		return
	}
	return
}

func checkErr(c *gin.Context, err error) {
	switch t := err.(type) {
	case *json.UnmarshalTypeError:
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(t.Field, httpcommon.RequestInvalid, t.Field))
		return
	case *json.SyntaxError:
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(err.Error(), httpcommon.RequestInvalid, ""))
		return
	case validator.ValidationErrors:
		errors := HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, httpcommon.NewPartialSuccess[any](false, nil, errors))
		return
	default:
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(err.Error(), httpcommon.RequestInvalid, ""))
		return
	}
}

func HandleValidationErrors(err error) (errors []httpcommon.ErrorResponse) {
	for _, fieldErr := range err.(validator.ValidationErrors) {
		message := ValidationErrorToText(fieldErr)
		fields := utils.LowerInitial(strings.Split(fieldErr.StructNamespace(), ".")[1:])
		field := strings.Join(fields, ".")
		errorCode, ok := TagMap[fieldErr.Tag()]
		if !ok {
			errorCode = httpcommon.InvalidFormat
		}
		errors = httpcommon.AddError(errors, message, string(errorCode), field)
	}
	return
}

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "gte", "max", "min":
		return constants.InvalidValue
	case "len":
		return constants.InvalidLength
	case "email":
		return constants.InvalidEmailFormat
	}
	return "InvalidRequest"
}

func GetQueryParamsHTTP(c *gin.Context, dest interface{}) (err error) {
	_, span := tracing.StartSpanFromContext(c.Request.Context(), "GetQueryParamsHTTP")
	defer span.End()
	if err = c.ShouldBindQuery(dest); err != nil {
		checkErr(c, err)
		return
	}
	return
}
