package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkerChan <- r
	}()

}

func (s *SimpleScheduler) ConfigureMasterWoekerChan(c chan engine.Request) {
	s.WorkerChan = c
}
