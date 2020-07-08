package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)">([^<]+)</a>`

// const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(b []byte) engine.ParseResult {
				return ParseProfile(b, "昵称："+name)
			},
		})
	}

	return result
}
