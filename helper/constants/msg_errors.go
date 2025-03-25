package constants

// Error code
const (
	SYSTEM_ERROR                    = "system_error"      //
	REQUEST_INVALID                 = "request_invalid"   //
	RECORD_NOT_EXIST                = "record_not_exist"  //
	RECORD_EXISTED                  = "record_existed"    //
	PERMISSION_DENIED               = "permission_denied" //
	UserIsNotExist                  = "USER_NOT_EXIST"
	InvalidFormat                   = "INVALID_FORMAT"
	ErrorMapData                    = "ERROR_MAP_DATA"
	CanNotFoundTheRoutes            = "CanNotFoundTheRoutes"
	StatusCDOrderInvalid            = "STATUS_CD_ORDER_INVALID"
	UpdatedAtIsChanged              = "UPDATEDAT_IS_CHANGED"
	DriverIsNotExist                = "DRIVER_NOT_EXIST"
	DriverIsNotActive               = "DRIVER_NOT_ACTIVE"
	ToKenIsMissing                  = "TOKEN_IS_MISSING"
	DuplicateType                   = "DUPLICATE_TYPE"
	ParentIdIsInvalid               = "ParentIdIsInvalid"
	WaypointsMustHaveThree          = "WaypointsMustHaveThree"
	UserIsNotAuthorized             = "USER_IS_NOT_AUTHORIZED"
	OrderIdIsInvalid                = "ORDER_ID_IS_INVALID"
	OrderHasBeenCanceled            = "ORDER_HAS_BEEN_CANCELED"
	AmountHasBeenChanged            = "AMOUNT_HAS_BEEN_CHANGED"
	PaymentAmountCanNotBeZero       = "PAYMENT_AMOUNT_CAN_NOT_BE_ZERO"
	UserHasNoPermissionToUseAutoPay = "USER_HAS_NO_PERMISSION_TO_USE_AUTO_PAY"
	GetUserFailed                   = "GET_USER_FAILED"
	InfoPaymentNotCorrect           = "INFO_PAYMENT_NOT_CORRECT"
	OrderHasNotBeenPaid             = "ORDER_HAS_NOT_BEEN_PAID"
	CustomerKeyNotMatch             = "CUSTOMER_KEY_NOT_MATCH"
	PerMissionDenied                = "PERMISSION_DENIED"
)

// Error message
const (
	SystemErrorMessage    = "There was an error on the server side"
	RequestInvalidMessage = "Invalid request"
	RecordNotExistMessage = "data does not exist"
	RecordExistMessage    = "data already exists"
)
