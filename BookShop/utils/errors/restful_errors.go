package errors

import "net/http"

// RestErr a struct to create & send errors witch are related to REST API
type RestErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

// NewBadRequestError returns a BAD_REQUEST 400
func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewNotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}
