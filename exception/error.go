package exception

type baseError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// --- / --- / --- / --- / --- / --- / --- / --- / ---

type UnexpectedError baseError

func NewUnexpectedError() error {
	return &UnexpectedError{
		Message: "Internal Server Error",
		Code:    "internal-server-error",
	}
}

func (b UnexpectedError) Error() string {
	return b.Message
}

// --- / --- / --- / --- / --- / --- / --- / --- / ---

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

// --- / --- / --- / --- / --- / --- / --- / --- / ---

type ItemNotFoundError baseError

func NewItemNotFoundError(code string, message string) error {
	return &ItemNotFoundError{
		Message: message,
		Code: code,
	}
}

func (b ItemNotFoundError) Error() string {
	return b.Message
}