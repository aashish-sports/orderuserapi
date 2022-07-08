package utils

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequest(messages string) *RestErr {
	return &RestErr{
		Message: messages,
		Status:  http.StatusInternalServerError,
		Error:   "bad Request",
	}
}
func NotFound(messages string) *RestErr {
	return &RestErr{
		Message: messages,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}
func NewInternalServerError(messages string) *RestErr {
	return &RestErr{
		Message: messages,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}
func AlreadyExistError(messages string) *RestErr {
	return &RestErr{
		Message: messages,
		Status:  http.StatusInternalServerError,
		Error:   "AlreadyExist",
	}
}
