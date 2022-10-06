package parser

import (
	"regexp"
)

var (
	cityListReg, cityReg, cityNextPageReg, userReg *regexp.Regexp
	err                                            error
)

type UserProfile struct {
	Name string
}

type Request struct {
	URL       string
	ParseFunc func([]byte) *ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []UserProfile
}

func init() {
	cityListReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]*)[^>]*>([^<]*)</a>`)
	cityReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]*)" target="_blank">`)
	cityNextPageReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/zhengzhou/\d*)">下一页</a>`)
	userReg = regexp.MustCompile(`<h1 data-v-cc1a17de="" class="nickName">(.*)</h1>`)
}

func ParseUser(bs []byte) *ParseResult {
	match := userReg.FindSubmatch(bs)
	user := UserProfile{
		Name: string(match[1]),
	}
	return &ParseResult{
		Requests: nil,
		Items:    []UserProfile{user},
	}
}

func ParseCity(bs []byte) *ParseResult {
	var requests []Request
	matches := cityReg.FindAllSubmatch(bs, -1)
	for _, m := range matches {
		url := m[1]
		requests = append(requests, Request{
			URL:       string(url),
			ParseFunc: ParseUser,
		})
	}

	// nextPageMatch := cityNextPageReg.FindSubmatch(bs)
	// requests = append(requests, Request{
	// 	URL:       string(nextPageMatch[1]),
	// 	ParseFunc: ParseCity,
	// })

	return &ParseResult{
		Requests: requests,
		Items:    nil,
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
