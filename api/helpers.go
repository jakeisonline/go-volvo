package api

import "fmt"

var marker string = "[go-volvo]"

func throwPanic(message string) {
	errorString := fmt.Sprintf("%s %s", marker, message)
	panic(errorString)
}

func makeError(message error) error {
	errorString := fmt.Errorf("%s %s", marker, message)
	return errorString
}
