package tickertape

import (
	"fmt"
	"strings"
	"time"
)

const (
	bufSize              = 30
	defaultPollingPeriod = 300 * time.Millisecond
	defaultSignalPeriod  = 1000 * time.Millisecond
)

type _signal struct{}

var signal _signal
var internalTicker TickerTape

//TickerTape is a struct for printing progress information one at a time, on the same line.
type TickerTape struct {
	listening                   bool
	events                      chan string
	signals                     chan _signal
	signalPeriod, pollingPeriod time.Duration
}

func (ticker *TickerTape) startListening() {
	if ticker.signalPeriod == 0 {
		ticker.signalPeriod = defaultSignalPeriod
	}
	if ticker.pollingPeriod == 0 {
		ticker.pollingPeriod = defaultPollingPeriod
	}
	ticker.listening = true
	ticker.events = make(chan string, bufSize)
	ticker.signals = make(chan _signal, bufSize)
	go ticker.listen()
	go ticker.repeatSignal()
}

func (ticker *TickerTape) repeatSignal() {

	for {
		ticker.signals <- signal
		time.Sleep(ticker.signalPeriod)
	}
}
func (ticker *TickerTape) listen() {
	for {
		select {
		case event := <-ticker.events:
			fmt.Print("\r", strings.Repeat(" ", 120))
			fmt.Print(strings.Repeat("\b", 120))
			fmt.Print("\r", event)
			// we want to be able to see each message as it comes up.
		case <-ticker.signals:
			fmt.Print(".")
		}
	}

}

//Printf prints to a ticker
func (ticker *TickerTape) Printf(format string, args ...interface{}) {
	if !ticker.listening {
		ticker.startListening()
	}
	ticker.events <- fmt.Sprintf(format, args...)
}

//Print a line through the ticker like fmt.Print
func (ticker *TickerTape) Print(args ...interface{}) {
	if !ticker.listening {
		ticker.startListening()
	}
	ticker.events <- fmt.Sprint(args...)
}

func (ticker *TickerTape) Write(b []byte) (n int, err error) {
	if !ticker.listening {
		return 0, fmt.Errorf("tickertape is not listening")
	}
	ticker.events <- string(b)
	return len(b), nil
}

//Printf to the internal ticker.
func Printf(format string, args ...interface{}) {
	internalTicker.Printf(format, args...)
}

//Print to the internal ticker.
func Print(args ...interface{}) {
	internalTicker.Print(args...)
}
