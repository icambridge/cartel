package cartel

import (
	"runtime"
	"sync"
	"time"
)

type Pool struct {
	Input  chan Task
	Output chan OutputValue
	wg     *sync.WaitGroup
}

func (p Pool) NumberOfItemsInQueue() int {
	return len(p.Input)
}

func (p Pool) End() {
	close(p.Input)
	p.wg.Wait()
}

func (p Pool) Do(t Task) {
	p.Input <- t
}

func (p Pool) GetOutput() []OutputValue {
	values := []OutputValue{}
	for {
		select {
		case r, ok := <-p.Output:
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

func (p Pool) worker() {
	t := time.Now()
	for {

		since := time.Since(t)

		if since.Minutes() > 2 {
			p.wg.Done()
			p.addWorker()
			break
		}

		t, ok := <-p.Input
		if !ok {
			p.wg.Done()
			// p.addWorker()
			break
		}
		v := t.Execute()
		p.Output <- v
		runtime.GC()
	}
}

func (p Pool) addWorker() {
	p.wg.Add(1)
	go p.worker()
}

func NewPool(numberOfWorkers int) Pool {

	jobs := make(chan Task, 100)
	results := make(chan OutputValue, 100)

	var wg sync.WaitGroup
	p := Pool{jobs, results, &wg}

	for w := 1; w <= numberOfWorkers; w++ {
		p.addWorker()
	}
	return p
}
