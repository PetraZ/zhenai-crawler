package persist

import (
	"log"

	"github.com/PetraZ/zhenai-crawler/model"
)

func NewItemSaver() chan model.UserProfile {
	c := make(chan model.UserProfile)
	go func() {
		itemCount := 0
		for {
			userProfile := <-c
			log.Printf("ItemSaver: got item %d %v", itemCount, userProfile)
			itemCount++
		}
	}()
	return c
}
