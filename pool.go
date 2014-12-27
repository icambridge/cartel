package cartel

type Pool struct {
	
}

type Task struct {
	
	
}

func worker(input <-chan Task, output chan<- string) {
	
	for {
		_, ok := <-input
		if !ok {
			break
		}
	}
}

func NewPool(numberOfWorkers int) Pool {
	
	jobs := make(chan Task, 100)
	results := make(chan string, 100)
	for w := 1; w <= numberOfWorkers; w++ {
		go worker(jobs, results)
	}
	return Pool{}
}
