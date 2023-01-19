package constant

type ErrorCode struct {
	Code string
	Msg  string
}

func NewErrorCode(code string, msg string) ErrorCode {
	return ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

var (
	OK                       = NewErrorCode("000000", "OK")
	Err                      = NewErrorCode("000001", "Fail")
	ParamErr                 = NewErrorCode("000002", "Parameters are empty or miss match")
	RequestReadErr           = NewErrorCode("000003", "network error")
	JsonStringToInterfaceErr = NewErrorCode("000004", "please read our API document")

	PasswordMismatch    = NewErrorCode("010000", "Please type the correct username or password")
	TokenGenerateFailed = NewErrorCode("010001", "Please type the correct username or password")
)
