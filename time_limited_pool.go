package cartel

import (
	"runtime"
	"sync"
	"time"
)

type TimeLimitedPool struct {
	input       chan Task
	output      chan interface{}
	wg          *sync.WaitGroup
	options     PoolOptions
	lastChecked time.Time
	allowace    float64
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
		p.rateLimit()
		t, ok := <-p.input

		if !ok {
			p.wg.Done()
			break
		}

		v := t.Execute()
		p.lastChecked = time.Now()

		p.output <- v
		runtime.GC()
	}
}

func (p *TimeLimitedPool) rateLimit() {
	lastChecked := p.lastChecked
	now := time.Now()

	p.lastChecked = now
	diff := float64(now.Unix() - lastChecked.Unix())
	floatRate := float64(p.options.PerDuration)
	floatDuration := float64(p.options.Duration.Seconds())

	p.allowace = p.allowace + (diff * (floatRate / floatDuration))
	if p.allowace < 1.0 {
		time.Sleep(p.options.Duration)
	}
}

func (p TimeLimitedPool) AddWorker() {
	p.wg.Add(1)
	go p.worker()
}

func NewTimeLimitedPool(options PoolOptions) TimeLimitedPool {

	jobs := make(chan Task, 100)
	results := make(chan interface{}, 100)

	var wg sync.WaitGroup
	p := TimeLimitedPool{jobs, results, &wg, options, time.Now(), 0.0}

	for w := 1; w <= options.Size; w++ {
		p.AddWorker()
	}
	return p
}
