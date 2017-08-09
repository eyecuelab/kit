package tickertape

import (
	"fmt"
	"time"
)

const (
	bufSize              = 30
	defaultPollingPeriod = 300 * time.Millisecond
	defaultSignalPeriod  = 500 * time.Millisecond
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

func repeatSignal(period time.Duration, ch chan<- _signal) {

}

func (ticker *TickerTape) startListening() {
	if ticker.signalPeriod == 0 {
		ticker.signalPeriod |= defaultSignalPeriod
	}
	if ticker.pollingPeriod == 0 {
		ticker.pollingPeriod = defaultPollingPeriod
	}
	ticker.listening = true
	go ticker.listen()
	go ticker.repeatSignal()
}

func (ticker *TickerTape) repeatSignal() {
	ticker.signals = make(chan _signal)
	for {
		ticker.signals <- signal
		time.Sleep(ticker.signalPeriod)
	}
}
func (ticker *TickerTape) listen() {
	ticker.events = make(chan string, bufSize)
	for i := 0; ; i++ {
		select {
		case event := <-ticker.events:
			fmt.Print("\r", event)
			time.Sleep(ticker.pollingPeriod) // we want to be able to see each message as it comes up.
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

//Printf to the internal ticker.
func Printf(format string, args ...interface{}) {
	internalTicker.Printf(format, args...)
}

//Print to the internal ticker.
func Print(args ...interface{}) {
	internalTicker.Print(args...)
}
