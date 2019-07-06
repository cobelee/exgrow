package errors

type AppError struct {
	Code     string
	ErrorMsg string
	Data     interface{}
}

func (e *AppError) Error() string {
	return e.ErrorMsg
}

// NewError 创建自定义Error对象
func NewError(code string, errorMsg string) AppError {
	var e AppError
	e.Code = code
	e.ErrorMsg = errorMsg
	return e
}
