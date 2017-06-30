package errors

import (
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

//FactoryChCheckF creates a partial function for chCheckF with the errors channel specified.
//MyChErrorF := FactorychCheck(errors)
//MyChErrorF(err, format, ...args) <-> ChErrorF(errors, err, format, ...args)
func FactoryChCheckF(errors chan<- error) func(err error, format string, args ...interface{}) {
	return func(err error, format string, args ...interface{}) {
		chCheckF(errors, err, format, args...)
	}
}

//Errors is the default error channel used by functions in this package.
var Errors = make(chan error)

//ChErrorF checks an error for nil and if non-nil, sends
//fmt.Errorf(format+": %v", args, err) to errors.Errors
var ChErrorF = FactoryChCheckF(Errors)
