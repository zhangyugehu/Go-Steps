package parser

import (
	"study/crawler/engine"
	"regexp"
)

var (
	profileReg = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	cityUrlReg = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {

	result := engine.ParseResult{}

	allMatches := profileReg.FindAllSubmatch(contents, -1)
	for _, match:=range allMatches {
		url := string(match[1])
		name := string(match[2])
		//result.Items = append(result.Items, "User: " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(match[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, url, name)
			},
		})
	}

	allMatches = cityUrlReg.FindAllSubmatch(contents, -1)
	for _, match:=range allMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(match[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
