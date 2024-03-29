package parser

import (
	"shishaoGo/crawler/crawler2/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 5
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			//ParserFunc: engine.NilParser,
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
