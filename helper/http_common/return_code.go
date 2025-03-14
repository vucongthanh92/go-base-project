package httpcommon

const (
	SystemError                     = "SYSTEM_ERROR"
	RequestInvalid                  = "REQUEST_INVALID"
	PickupTimeInvalid               = "PICKUP_TIME_INVALID"
	IsExist                         = "EXISTED"
	IsNotExist                      = "NOT_EXIST"
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

const (
	OtpMessage                = "OTP is not correct"
	MappingObjectError        = "Mapping Object Error"
	OrderIdNotCorrect         = "OrderId Not Correct"
	UpdatedAtIsChangedMessage = "Data has been changed"
)
