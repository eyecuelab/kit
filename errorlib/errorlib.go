package errorlib

import (
	"errors"
	"fmt"
	"log"
)

//chErrorF checks an error for nil and if non-nil, sends
//fmt.Errorf(format+": %v", args, err) to the channel specified.
//Note that this is deliberately unexported.
func chCheckF(errors chan<- error, err error, format string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		format += ": %v"
		errors <- fmt.Errorf(format, args...)
	}
}

//FatalCheckF checks an error for nil and calls log.Fatalf(msg+": %v", args, err) if non-nil.
func FatalCheckF(err error, format string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		format += ": %v"
		log.Fatalf(format, args...)
	}
}

//FactoryChCheckF creates a partial function for chCheckF with the errors channel specified. Equivalent to
//MyChErrorF := FactorychCheck(errors); MyChErrorF(err, format, ...args) <-> ChErrorF(errors, err, format, ...args)
func FactoryChCheckF(errors chan<- error) func(err error, format string, args ...interface{}) {
	return func(err error, format string, args ...interface{}) {
		chCheckF(errors, err, format, args...)
	}
}

//NewChecker produces an error channel, starts LogErrors() on that channel in it's own goroutine,
//and produces a checker function that sends non-fatal errors to that channel. Equivalent to
//errors := LoggedChannel(); //checker := FactoryChCheckF(errors)
func NewChecker() (errors chan error, checker func(err error, format string, args ...interface{})) {
	errors = make(chan error)
	go LogErrors(errors)
	return errors, FactoryChCheckF(errors)
}

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
		fmt.Println(err)
	}
}

//Errors is the default error channel used by functions in this package.
var Errors = make(chan error)

//ChErrorF checks an error for nil and if non-nil, sends
//fmt.Errorf(format+": %v", args, err) to errors.Errors
var ChErrorF = FactoryChCheckF(Errors)

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
