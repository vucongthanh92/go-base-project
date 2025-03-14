package validation

type TagErrorCode map[string]ErrorCode

type ErrorCode string

const (
	InvalidErrorCode ErrorCode = "invalid"
)

var TagMap TagErrorCode

func init() {
	TagMap = map[string]ErrorCode{}
}
