package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    http.StatusBadRequest,
		Error:   "Bad Request",
	}
}
