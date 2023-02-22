package main

import (
	"shishaoGo/crawler/crawler6/engine"
	"shishaoGo/crawler/crawler6/scheduler"
	"shishaoGo/crawler/crawler6/zhenai/parser"
)

func main() {
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{}, // Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
