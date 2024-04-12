package api

import "fmt"

var marker string = "[go-volvo]"

func ThrowPanic(message string) {
	errorString := fmt.Sprintf("%s %s", marker, message)
	panic(errorString)
}

func MakeError(message error) error {
	errorString := fmt.Errorf("%s %s", marker, message)
	return errorString
}
