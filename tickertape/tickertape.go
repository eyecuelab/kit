//Package tickertape provides an implenetation of a concurrency-safe 'ticker tape' of current information during  a running program - that is, repeatedly updating the same line with new information.
package tickertape

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/eyecuelab/kit/errorlib"
)

const (
	bufSize              = 30
	defaultPollingPeriod = 300 * time.Millisecond
	defaultSignalPeriod  = 1000 * time.Millisecond

	ErrNotListening  errorlib.ErrorString = "tickertape is not listening"
	ErrAlreadyClosed errorlib.ErrorString = "cannot print to a closed tickertape"
)

type _signal struct{}

var signal _signal
var internalTicker TickerTape

var _ io.WriteCloser = &TickerTape{}

//TickerTape is a struct for printing progress information one at a time, on the same line.
type TickerTape struct {
	mux       sync.Mutex //guards listening and closed
	listening bool
	closed    bool

	done    chan _signal
	events  chan string
	signals chan _signal

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
	ticker.done = make(chan _signal, 2)
	go ticker.listen()
	go ticker.repeatSignal()
}

func (ticker *TickerTape) repeatSignal() {
	defer close(ticker.signals)
	for {
		select {
		case <-ticker.done:
			return

		default:
			ticker.signals <- signal
			time.Sleep(ticker.signalPeriod)
		}
	}
}

//Close the tickertape.
func (ticker *TickerTape) Close() error {
	defer close(ticker.done)
	ticker.done <- signal // stop repeatSignal
	ticker.done <- signal // stop listen
	return nil
}

func (ticker *TickerTape) listen() {
	defer close(ticker.events)
	for {
		select {
		case event := <-ticker.events:
			fmt.Print("\r", strings.Repeat(" ", 120))
			fmt.Print(strings.Repeat("\b", 120))
			fmt.Print("\r", event)
			// we want to be able to see each message as it comes up.
		case <-ticker.signals:
			fmt.Print(".")
		case <-ticker.done:
			close(ticker.events)
			return
		}
	}

}

//Printf prints to a ticker
func (ticker *TickerTape) Printf(format string, args ...interface{}) bool {
	return ticker.print(fmt.Sprintf(format, args...))
}

func (ticker *TickerTape) print(s string) bool {
	ticker.mux.Lock()
	if ticker.closed {
		return false
	}

	if !ticker.listening {
		ticker.startListening()
	}
	ticker.mux.Unlock()
	ticker.events <- s
	return true
}

//Print a line through the ticker like fmt.Print
func (ticker *TickerTape) Print(args ...interface{}) bool {
	return ticker.print(fmt.Sprint(args...))
}

func (ticker *TickerTape) Write(b []byte) (n int, err error) {
	if !ticker.listening {
		return 0, ErrNotListening
	}
	ok := ticker.print(string(b))
	if !ok {
		return 0, ErrAlreadyClosed
	}
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
