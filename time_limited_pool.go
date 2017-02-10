package cartel

import (
	"sync"
	"time"
)

type TimeLimitedPool struct {
	input     chan Task
	output  chan interface{}
	wg      *sync.WaitGroup
	available map[int]bool
	options PoolOptions
}


func (p TimeLimitedPool) End() {
	close(p.input)
	p.wg.Wait()
}

func (p TimeLimitedPool) Do(t Task) {
	p.input <- t
}

func (p TimeLimitedPool) GetOutput() []interface{} {
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

func (p *TimeLimitedPool) worker() {
	for {
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

func (p *TimeLimitedPool) AddWorker() {
	p.wg.Add(1)
	go p.worker()
}
func NewTimeLimitedPool(options PoolOptions) TimeLimitedPool {

	jobs := make(chan Task, 100)
	results := make(chan interface{}, 100)
	available := map[int]bool{}

	var wg sync.WaitGroup
	p := TimeLimitedPool{jobs, results, &wg, available, options}

	for w := 1; w <= options.Size; w++ {
		p.AddWorker()
	}
	return p
}