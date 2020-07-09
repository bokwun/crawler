package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}
//每次调用WorkerChan()，都会返回已经创建好的channel（唯一的channel）
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}