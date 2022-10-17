package engine

import (
	"log"

	"github.com/PetraZ/zhenai-crawler/fetcher"
	"github.com/PetraZ/zhenai-crawler/parser"
)

func HandleRequest(r parser.Request) (*parser.ParseResult, error) {
	log.Printf("Fetching %s", r.URL)
	bs, err := fetcher.Fetch(r.URL)
	if err != nil {
		return nil, err
	}
	parserResult := r.ParseFunc(bs)
	return parserResult, nil
}
