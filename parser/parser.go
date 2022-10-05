package parser

import (
	"regexp"
)

var (
	cityListReg *regexp.Regexp
	err         error
)

type Request struct {
	URL       string
	ParseFunc func([]byte) *ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func init() {
	cityListReg, err = regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[^"]*)[^>]*>([^<]*)</a>`)
	if err != nil {
		panic(err)
	}
}

func ParseCityList(bs []byte) *ParseResult {
	var requests []Request
	matches := cityListReg.FindAllSubmatch(bs, -1)
	for _, m := range matches {
		url := m[1]
		requests = append(requests, Request{
			URL:       string(url),
			ParseFunc: ParseCity})
	}
	return &ParseResult{
		Requests: requests,
		Items:    nil,
	}
}

func ParseCity(bs []byte) *ParseResult {
	return nil
}
