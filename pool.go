package cartel

import (
	"runtime"
	"sync"
	"time"
)

type PoolOptions struct {
	PerDuration int
	Duration    time.Duration
	Size        int
}

type Pool interface {
	Do(t Task)
	GetOutput() []interface{}
	End()
}

type NoTimeLimitPool struct {
	input   chan Task
	output  chan interface{}
	wg      *sync.WaitGroup
}

func (p NoTimeLimitPool) End() {
	close(p.input)
	p.wg.Wait()
}

func (p NoTimeLimitPool) Do(t Task) {
	p.input <- t
}

func (p NoTimeLimitPool) GetOutput() []interface{} {
	values := []interface{}{}
	for {
		select {
		case r, ok := <-p.output:
			if ok {
				values = append(values, r)
			} else {
				return values
			}
		default:
			return values
		}
	}
}

func (p *NoTimeLimitPool) worker() {
	t := time.Now()
	for {

		since := time.Since(t)

		if since.Minutes() > 2 {
			p.wg.Done()
			p.AddWorker()
			break
		}

		t, ok := <-p.input
		if !ok {
			p.wg.Done()
			break
		}
		v := t.Execute()
		p.output <- v
		runtime.GC()
	}
}

func (p *NoTimeLimitPool) AddWorker() {
	p.wg.Add(1)
	go p.worker()
}

func NewPool(options PoolOptions) Pool {

	jobs := make(chan Task, 100)
	results := make(chan interface{}, 100)

	var wg sync.WaitGroup
	p := NoTimeLimitPool{jobs, results, &wg}

	for w := 1; w <= options.Size; w++ {
		p.AddWorker()
	}
	return p
}
