package log

import (
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
	log.Infof(format, args)
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	if goenv.Prod {
		log.SetFormatter(&log.JSONFormatter{})
	}
}
