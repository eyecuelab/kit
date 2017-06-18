package log

import (
	log "github.com/sirupsen/logrus"
)

func Fatal(err error) {
	log.Fatalf("%+v", err)
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}
