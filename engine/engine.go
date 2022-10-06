package engine

import (
	"log"

	"github.com/PetraZ/zhenai-crawler/fetcher"
	"github.com/PetraZ/zhenai-crawler/parser"
)

func Run(seeds []parser.Request) error {
	var requests []parser.Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		bs, err := fetcher.Fetch(request.URL)
		log.Printf("Fetching %s", request.URL)
		if err != nil {
			return err
		}
		parserResult := request.ParseFunc(bs)
		if parserResult == nil {
			continue
		}
		if parserResult.Items != nil {
			log.Print(parserResult.Items[0])
		}
		for _, newRequest := range parserResult.Requests {
			requests = append(requests, newRequest)
		}
	}
	return nil
}
