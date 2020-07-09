package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// 每次调用WorkerChan()，都会创建一个channel
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// 哪个worker准备好了，就把它的chan传给workerChan通道
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			// 定义的activeWorker为nil channel
			var activeWorker chan engine.Request
			// 必要，让两者都有数据在执行第三个case，不然有channel，只会传空值给channel
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				//实例化activeWorker这个channel，这样第三个case就能被选到
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			// activeWorker若为nil channel，则永远select不到这个case
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
