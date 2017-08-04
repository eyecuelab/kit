package clog

import (
	"fmt"
	"strings"
)

type Logger chan string

func (log Logger) Flog(format string, args ...interface{}) {
	log <- fmt.Sprintf(format, args...)
}

func LoggedChannel() Logger {
	logCH := make(chan string)
	go LogChannel(logCH)
	return logCH
}

func LogChannel(log Logger) {
	var prevLine string
	for line := range log {
		fmt.Print(strings.Repeat("\b", len(prevLine)))
		fmt.Print(line)
		prevLine = line
	}
}
