package main

import (
	"crawler/businessName/parser"
	"crawler/engine"
	"crawler/scheduler"
)

func main() {
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
