package parser

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

const profileRe = `<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`

func ParseProfile(contents []byte) engine.ParseResult {
	fmt.Println("test!")
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "Profile "+string(m[1]))
	}
	return result
}
