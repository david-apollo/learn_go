package engine

import (
	"log"
)

// ConcurrentEngine type
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan  chan Item
}


// Scheduler interface
type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// ReadNotifier interface
type ReadNotifier interface {
	WorkerReady(chan Request)
}


// Run func for ConcurrentEngine 
func (e ConcurrentEngine) Run(seeds ...Request) {

	// in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParserResult, ready ReadNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}


var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
