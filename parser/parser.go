package parser

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// <tr>
// <td width="\d*">
// <span class="grayL">年龄：</span>"([^<]{1,10})"
// </td>
// </tr>

var (
	// On city list page, get the city list
	cityListReg     = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]*)[^>]*>([^<]*)</a>`)
	cityNextPageReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/zhengzhou/\d*)">下一页</a>`)
)

type Request struct {
	URL       string
	ParseFunc func([]byte) *ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []UserProfile
}

// Top level - city list
func ParseCityList(bs []byte) *ParseResult {
	var requests []Request
	matches := cityListReg.FindAllSubmatch(bs, -1)
	for _, m := range matches {
		url := m[1]
		requests = append(requests, Request{
			URL:       string(url),
			ParseFunc: ParseCity})
	}
	fmt.Println("The number of cities are: ", len(matches))
	return &ParseResult{
		Requests: requests,
		Items:    nil,
	}
}

type UserProfile struct {
	ID           string
	Name         string
	URL          string
	Sex          string
	Location     string
	Age          string
	Education    string
	Marriage     string
	Height       string
	Salary       string
	Introduction string
}

func ParseCity(bs []byte) *ParseResult {
	var requests []Request
	var users []UserProfile
	// Parsing html page
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".list-item > .content").Each(func(i int, s *goquery.Selection) {
		user := UserProfile{}
		user.Name = s.Find("a").Text()
		personUrl, _ := s.Find("a").Attr("href")
		user.URL = personUrl

		s.Find("td").Each(func(i int, sub *goquery.Selection) {
			kv := strings.Split(sub.Text(), `：`)
			if len(kv) < 2 {
				return
			}
			k := strings.ReplaceAll(kv[0], "\u00a0", " ")
			switch k {
			case "性别":
				user.Sex = kv[1]
			case "居住地":
				user.Location = kv[1]
			case "年龄":
				user.Age = kv[1]
			case "学   历":
				user.Education = kv[1]
			case "婚况":
				user.Marriage = kv[1]
			case "身   高":
				user.Marriage = kv[1]
			case "月   薪":
				user.Salary = kv[1]
			default:
				fmt.Printf("User profile kv is not recognized k: %s v: %s\n", kv[0], kv[1])
			}

		})
		user.Introduction = s.Find(".introduce").Text()
		users = append(users, user)
		// fmt.Printf("%v, %v \n", i, user)
	})

	// nextPageMatch := cityNextPageReg.FindSubmatch(bs)
	// requests = append(requests, Request{
	// 	URL:       string(nextPageMatch[1]),
	// 	ParseFunc: ParseCity,
	// })

	return &ParseResult{
		Requests: requests,
		Items:    users,
	}
}
