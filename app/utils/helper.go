package utils

import (
	// "math/rand"
	// "time"
	"regexp"
	"strconv"
)

type ErrorType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomHTTPError struct {
	Error ErrorType `json:"error"`
}

type CustomHTTPResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ValidateEmail(email string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(email)
	return
}

func GenerateError(errorCode int, msg string) (_error CustomHTTPError) {
	_error = CustomHTTPError{
		Error: ErrorType{
			Code:    errorCode,
			Message: msg,
		},
	}
	return
}

func ConvertToUint(str string) uint {
	uintNumber, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(uintNumber)

}
