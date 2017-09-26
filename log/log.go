package log

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/eyecuelab/kit/goenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	//Print is an alias for log.Print
	Print = log.Print

	//Fatalf is an alias for log.Fatalf
	Fatalf = log.Fatalf

	//Fatal is an alias for log.Fatal
	Fatal = log.Fatal
)

func FatalWrap(err error, msg string) {
	log.Fatalf("%+v", errors.Wrap(err, msg))
}

//Info is an alias for log.Info in the standard library
func Info(msg string) {
	log.Info(msg)
}

func Infof(format string, args ...interface{}) {
	s := spew.Sprintf(format, args...)
	log.Info(s)
}

func Infofln(format string, args ...interface{}) {
	fmt.Println("")
	Infof(format, args...)
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
