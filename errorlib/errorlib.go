package errorlib

import (
	"errors"
	"fmt"

	"github.com/eyecuelab/kit/log"
)

//ErrorString is a string with an Error() method. This lets you declare errors as compile-time constants,
//which facilitates various IDE tools.
type ErrorString string

func (err ErrorString) Error() string {
	return string(err)
}

const someErr = ErrorString("")

//LoggedChannel creates an error channel and starts LogErrors on that channel in it's own goroutine.
func LoggedChannel() chan error {
	errors := make(chan error)
	go LogErrors(errors)
	return errors
}

//LogErrors logs errors as they come in to os.stdout.
//Goroutine. Do not run directly!
func LogErrors(errors <-chan error) {
	for err := range errors {
		log.Print(err)
	}
}

//Errors is the default error channel used by functions in this package.
var Errors = make(chan error)

//Flatten a slice of errors into a single error
func Flatten(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	errmsg := "one or more errors: "
	for _, err := range errs {
		errmsg += fmt.Sprintf("%v; ", err)
	}
	return errors.New(errmsg)
}
