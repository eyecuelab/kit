package log

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/eyecuelab/kit/goenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	//Print is an alias for log.Print in logrus
	Print = log.Print

	//Fatalf is an alias for log.Fatalf in logrus
	Fatalf = log.Fatalf

	//Fatal is an alias for log.Fatal in logrus
	Fatal = log.Fatal

	//Println is an alias for log.Println in logrus
	Println = log.Println

	//Info is an alias for log.Info in logrus
	Info = log.Info

	//Warn is an alias for log.Warn in logrus
	Warn = log.Warn

	//Warnf is an alias for log.Warnf in logrus
	Warnf = log.Warnf
)

func FatalWrap(err error, msg string) {
	log.Fatalf("%+v", errors.Wrap(err, msg))
}

func Infof(format string, args ...interface{}) {
	s := spew.Sprintf(format, args...)
	log.Infof(s)
}

func Infofln(format string, args ...interface{}) {

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
