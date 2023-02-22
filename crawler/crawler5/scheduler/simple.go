package scheduler

import (
	"shishaoGo/crawler/crawler5/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

/*func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}*/

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send requst down to worker chan
	//s.workerChan <- r // 实现1，会卡死

	// 实现2：并发分发Request
	go func() {
		s.workerChan <- r
	}()
}
