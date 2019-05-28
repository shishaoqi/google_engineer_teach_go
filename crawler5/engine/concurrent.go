package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request) // 接收Request
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request) //拿出去创建 ReadyNotifier
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item : %v total = %d", item, itemCount)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
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