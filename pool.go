package cartel

import (
	"sync"
)
type Pool struct {
	jobs chan Task
	wg *sync.WaitGroup
	
}

func (p Pool) End() {
	close(p.jobs)
	p.wg.Wait()
}

type Task struct {
	
	
}

func (p Pool) worker(input <-chan Task, output chan<- string) {

	for {
		_, ok := <-input
		if !ok {
			p.wg.Done()
			break
		}
	}
}

func NewPool(numberOfWorkers int) Pool {


	jobs := make(chan Task, 100)
	results := make(chan string, 100)

	var wg sync.WaitGroup
	p := Pool{jobs, &wg}
	
	for w := 1; w <= numberOfWorkers; w++ {
		wg.Add(1)
		go p.worker(jobs, results)
	}
	return p
}
