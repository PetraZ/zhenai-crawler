package engine

import (
	"log"

	"github.com/PetraZ/zhenai-crawler/model"
	"github.com/PetraZ/zhenai-crawler/parser"
)

type SimpleEngine struct{}

func (SimpleEngine) Run(seeds []parser.Request) error {
	var requests []parser.Request

	requests = append(requests, seeds...)

	var users []model.UserProfile
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parserResult, err := HandleRequest(request)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if parserResult == nil {
			continue
		}
		if parserResult.Items != nil {
			users = append(users, parserResult.Items...)
			log.Printf("We now have %v user profiles", len(users))
		}
		if parserResult.Requests != nil {
			requests = append(requests, parserResult.Requests...)
		}
	}
	return nil
}
