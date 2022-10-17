package main

import (
	"github.com/PetraZ/zhenai-crawler/engine"
	"github.com/PetraZ/zhenai-crawler/parser"
)

func main() {
	engine.ConcurrentEngine{}.Run([]parser.Request{
		{
			URL:       "http://www.zhenai.com/zhenghun",
			ParseFunc: parser.ParseCityList,
		},
	})
	// r, err := fetcher.Fetch("https://www.zhenai.com/zhenghun/xiamen")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// parser.ParseCity(r)
}
