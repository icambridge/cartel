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
	AddWorker()
	Do(t Task)
	GetOutput() []interface{}
	End()
}

type NonTimeLimitedPool struct {
	input   chan Task
	output  chan interface{}
	wg      *sync.WaitGroup
}

func (p NonTimeLimitedPool) End() {
	close(p.input)
	p.wg.Wait()
}

func (p NonTimeLimitedPool) Do(t Task) {
	p.input <- t
}

func (p NonTimeLimitedPool) GetOutput() []interface{} {
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

func (p *NonTimeLimitedPool) worker() {
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

func (p *NonTimeLimitedPool) AddWorker() {
	p.wg.Add(1)
	go p.worker()
}

func NewPool(options PoolOptions) Pool {
	return NewNonTimeLimitedPool(options)
}

func NewNonTimeLimitedPool(options PoolOptions) NonTimeLimitedPool {

	jobs := make(chan Task, 100)
	results := make(chan interface{}, 100)

	var wg sync.WaitGroup
	p := NonTimeLimitedPool{jobs, results, &wg}

	for w := 1; w <= options.Size; w++ {
		p.AddWorker()
	}
	return p
}