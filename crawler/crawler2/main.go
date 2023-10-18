package main

import (
	"shishaoGo/crawler/crawler2/engine"
	"shishaoGo/crawler/crawler2/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
