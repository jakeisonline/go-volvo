package helpers

import "fmt"

var marker string = "[go-volvo]"

func Panic(message string) {
	errorString := fmt.Sprintf("%s %s", marker, message)
	panic(errorString)
}

func Error(message error) error {
	errorString := fmt.Errorf("%s %s", marker, message)
	return errorString
}
