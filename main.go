package main

import (
	"github.com/PetraZ/zhenai-crawler/engine"
	"github.com/PetraZ/zhenai-crawler/parser"
	"github.com/PetraZ/zhenai-crawler/persist"
)

func main() {
	engine.ConcurrentEngine{
		NumWorkers:   10,
		ItemSaveChan: persist.NewItemSaver(),
	}.Run(
		[]parser.Request{
			{
				URL:       "http://www.zhenai.com/zhenghun",
				ParseFunc: parser.ParseCityList,
			},
		})
}
