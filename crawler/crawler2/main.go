package main

import (
	"github.com/shishao/hello/crawler/crawler2/engine"
	"github.com/shishao/hello/crawler/crawler2/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}