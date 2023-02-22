package main

import (
	"shishaoGo/crawler/crawler5/engine"
	"shishaoGo/crawler/crawler5/scheduler"
	"shishaoGo/crawler/crawler5/zhenai/parser"
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
