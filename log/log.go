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

func Fatal(err error) {
	log.Fatal(err)
}

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

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var Logger = log.StandardLogger()

func init() {
	if goenv.Prod {
		log.SetFormatter(&log.JSONFormatter{})
	}
}
