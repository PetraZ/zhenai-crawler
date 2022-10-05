package main

import (
	"github.com/PetraZ/zhenai-crawler/engine"
	"github.com/PetraZ/zhenai-crawler/parser"
)

func main() {
	engine.Run([]parser.Request{
		{
			URL:       "http://www.zhenai.com/zhenghun",
			ParseFunc: parser.ParseCityList,
		},
	})
}
