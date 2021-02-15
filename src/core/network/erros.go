package network

import (
	"fmt"
)

type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

func NewApiError(code int64, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func (e APIError) Error() string {
	return fmt.Sprintf("[APIError] code=%d, msg=%s", e.Code, e.Message)
}

func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
