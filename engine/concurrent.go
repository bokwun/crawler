package engine

import "log"

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	//main.go 里queued实例化了一个接口类型，所以这个Run是queued的
	e.Scheduler.Run()
	//创建一定数量的worker协程
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 分配任务
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 接受worker完成返回的信息
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d %v\n", itemCount, item)
			itemCount++
		}
		//解析后再把新的任务分配给新的worker
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//创建好worker之后，把对应的chan加入队列
			ready.WorkerReady(in)
			//接收到request
			request := <-in
			//开始工作
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
