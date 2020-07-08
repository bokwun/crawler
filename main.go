package main

import (
	"crawler/businessName/parser"
	"crawler/engine"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

// const text = `
// my email is chenboqun@protonmail.com
// email1 is 1231313@qq.com
// email2 is 99898@163.com
// email3 is 098129389@qq.com.cn
// `

// func main() {
// 	re := regexp.MustCompile(`([a-zA-z0-9]+)@([a-zA-z0-9]+)(\.[a-zA-z0-9.]+)`)
// 	// match := re.FindAllString(text, -1)
// 	match := re.FindAllStringSubmatch(text, -1)
// 	for _, m := range match {
// 		for _, m1 := range m {
// 			fmt.Println("   ", m1)
// 		}
// 	}
// }
