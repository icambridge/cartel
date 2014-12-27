package cartel

import (
	"sync"
)
type Pool struct {
	Input chan Task
	Output chan OutputValue
	wg *sync.WaitGroup
}

func (p Pool) End() {
	close(p.Input)
	p.wg.Wait()
}

func (p Pool) Do(t Task) {
	p.Input <- t
}

func (p Pool) worker() {

	for {
		t, ok := <-p.Input
		if !ok {
			p.wg.Done()
			break
		}
		v := t.Execute()
		p.Output <- v
	}
}

func NewPool(numberOfWorkers int) Pool {


	jobs := make(chan Task, 100)
	results := make(chan OutputValue, 100)

	var wg sync.WaitGroup
	p := Pool{jobs, results, &wg}
	
	for w := 1; w <= numberOfWorkers; w++ {
		wg.Add(1)
		go p.worker()
	}
	return p
}
