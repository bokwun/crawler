package parser

import (
	"crawler/engine"
	"regexp"
)

const profileRe = `<td><span class="label">([^<]+)</span>([^<]+)</td>`

// const profileRe = `<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	item := []string{"Profile:", name}
	for _, m := range matches {
		item = append(item, string(m[1])+string(m[2]))
	}
	result.Items = append(result.Items, item)
	return result
}
