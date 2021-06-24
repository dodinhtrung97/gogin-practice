package error

type HttpError struct {
	Message string
	Status  int
	Error   string
}

func NewNotFoundError(message string) HttpError {
	return HttpError{
		Message: message,
		Status:  404,
		Error:   "Not Found",
	}
}

func NewBadRequestError(message string) HttpError {
	return HttpError{
		Message: message,
		Status:  400,
		Error:   "Bad Request",
	}
}
