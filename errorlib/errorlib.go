//Package errorlib contains tools to deal with errors, including the heavily-used `ErrorString` type (for errors that are compile-time constants) and `LoggedChannel` function (to report non-fatal errors during concurrent execution)
package errorlib

import (
	"errors"

	"github.com/eyecuelab/kit/log"
)

//ErrorString is a string with an Error() method. This lets you declare errors as visible compile-time constants,
//which facilitates various IDE tools.
type ErrorString string

func (err ErrorString) Error() string {
	return string(err)
}

const someErr ErrorString = "foo"

var _ error = someErr //satisfies interface

//LoggedChannel creates an error channel that will automatically log errors sent to it to the logger specified in kit/log.
func LoggedChannel() chan error {
	errors := make(chan error)
	go logErrors(errors)
	return errors
}

//logErrors logs errors as they come in to os.stdout.
//Goroutine. Do not run directly!
func logErrors(errors <-chan error) {
	for err := range errors {
		log.Print(err)
	}
}

//Errors is the default error channel used by functions in this package.
var Errors = make(chan error)

//Flatten a slice of errors into a single error
func Flatten(errs []error) error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		errmsg := "multiple errors:"
		for _, err := range errs {
			errmsg += err.Error()
		}
		return errors.New(errmsg)
	}
}
