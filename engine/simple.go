package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	// 任务队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		//队列里第一个任务，里面包含 url 和 对应的解析器
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		//补充队列任务
		requests = append(requests, parseResult.Requests...)
		//打印item 城市列表信息
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	// 拉取信息
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	// 解析信息
	return r.ParserFunc(body), nil
}
