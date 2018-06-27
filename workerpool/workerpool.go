package workerpool

import (
	"sync"
	"github.com/eyecuelab/kit/brake"
	"sync/atomic"
	"fmt"
)

const (
	statusMsg = "\r%v/%v (%.3f%%) | Success: %v  Errored: %v               "
)

type Task struct {
	Err error

	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}

func (t *Task) run() error {
	t.Err = t.f()
	return t.Err
}

type Pool struct {
	TasksChan  chan *Task
	ErrorsChan chan error

	concurrency  uint
	taskWg       sync.WaitGroup
	errorWg      sync.WaitGroup
	errorCount   uint64
	successCount uint64
}

func NewPool(concurrency uint) *Pool {
	return &Pool{
		TasksChan:   make(chan *Task),
		concurrency: concurrency,
		ErrorsChan:  make(chan error),
	}
}

func (p *Pool) Run() {
	for i := uint(0); i < p.concurrency; i++ {
		p.taskWg.Add(1)
		go p.work()
	}

	brake.NotifyFromChan(p.ErrorsChan, &p.errorWg)

	p.taskWg.Wait()

	close(p.ErrorsChan)
	p.errorWg.Wait()
}

func (p *Pool) PrintStatus(total int) {
	progressCount := p.successCount + p.errorCount
	percentDone := float64(progressCount) / float64(total) * 100
	fmt.Printf(statusMsg, progressCount, total, percentDone, p.successCount, p.errorCount)
}

func (p *Pool) work() {
	for task := range p.TasksChan {
		if err := task.run(); err != nil {
			p.ErrorsChan <- err
			atomic.AddUint64(&p.errorCount, 1)
			continue
		}
		atomic.AddUint64(&p.successCount, 1)
	}
	p.taskWg.Done()
}
