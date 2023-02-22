package main

import (
	"shishaoGo/crawler/crawler4/engine"
	"shishaoGo/crawler/crawler4/scheduler"
	"shishaoGo/crawler/crawler4/zhenai/parser"
)

func main() {
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
