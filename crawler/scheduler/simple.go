package scheduler

import "learn_go/crawler/engine"

// SimpleScheduler type
type SimpleScheduler struct {
	workerChan chan engine.Request
}


// WorkerChan func
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// Submit func
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

// Run func
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}