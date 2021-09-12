package exception

type baseError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type BadRequestError baseError

func NewBadRequestError(code string, message string) error {
	return &BadRequestError{
		Message: message,
		Code:    code,
	}
}

func (b BadRequestError) Error() string {
	return b.Message
}