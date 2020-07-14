package parser

import (
	"../../engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User " + name)
		result.Requests = append(
			result.Requests, engine.Request{
			Url:        string(m[1]),
			//ParserFunc: engine.NilParser,
			//ParserFunc: PaseProfile,
			ParserFunc: func(c []byte) engine.ParseResult {
				return PaseProfile(c, name)
			},
		})
	}

	return result
}
