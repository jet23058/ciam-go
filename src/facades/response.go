package facades

type ApiError struct {
	StatusCode int
	ErrorType  ErrorType
	Error      error
}

type ApiSuccess struct {
	StatusCode int
	Data       interface{}
}

type ResponseSuccess struct {
	Data string `example:"success"`
}
