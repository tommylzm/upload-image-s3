package errorcode

import "net/http"

const (
	DbLinkError       = "1101" // DB連線錯誤
	RedisLinkError    = "1102" // REDIS連線錯誤
	OtherServiceError = "1103" // 呼叫其他服務錯誤
	ExceptionError    = "1104" // 例外錯誤
	InputInvalid      = "1105" // 參數錯誤
	SqlError          = "1006"
)

// CustomError :
type CustomError struct {
	HTTPStatus int    `json:"HTTPStatus"`
	Code       string `json:"Code"`
	Message    string `json:"Message"`
}

func (c *CustomError) Error() string {
	return c.Message
}

func (c *CustomError) Decode() (int, string, string) {
	return c.HTTPStatus, c.Code, c.Message
}

// InputInvalidError : 參數錯誤
func InputInvalidError(err error) error {
	return &CustomError{
		HTTPStatus: http.StatusBadRequest,
		Code:       InputInvalid,
		Message:    err.Error(),
	}
}
