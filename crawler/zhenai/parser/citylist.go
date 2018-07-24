package parser

import (
	"study/crawler/engine"
	"regexp"
)

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]*)</a>`)
	allMatches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, match:=range allMatches {
		//result.Items = append(result.Items, "City: " + string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(match[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
