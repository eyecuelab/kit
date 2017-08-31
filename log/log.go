package log

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/eyecuelab/kit/goenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func FatalWrap(err error, msg string) {
	log.Fatalf("%+v", errors.Wrap(err, msg))
}

//Fatal is an alias for log.Fatal
func Fatal(err error) {
	log.Fatal(err)
}

//Fatalf is an alias for log.Fatalf in the standard library
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(err, args...)
}

//Info is an alias for log.Info in the standard library
func Info(msg string) {
	log.Info(msg)
}

func Infof(format string, args ...interface{}) {
	s := spew.Sprintf(format, args...)
	log.Info(s)
}

func ErrorWrap(err error, msg string) {
	log.Errorf("%+v", errors.Wrap(err, msg))
}

//Check calls Fatal on an error if it is non-nil.
func Check(err error) {
	if err != nil {
		Fatal(err)
	}
}

var Logger = log.StandardLogger()

func init() {
	if goenv.Prod {
		log.SetFormatter(&log.JSONFormatter{})
	}
}
