package scheduler

import "learn_go/crawler/engine"

// SimpleScheduler type
type SimpleScheduler struct {
	workerChan chan engine.Request
}


// ConfigureMasterWorkerChan func Scheduler
func (s SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

// Submit func
func (s SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}