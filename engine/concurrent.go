package engine

import "log"

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWoekerChan(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWoekerChan(in)
	//创建一定数量的worker协程
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
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

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result

		}
	}()
}
