package engine

import (
	"log"

	"github.com/PetraZ/zhenai-crawler/parser"
)

type ConcurrentEngine struct{}

func (e ConcurrentEngine) Run(seeds []parser.Request) error {

	// Creating work pool
	in := make(chan parser.Request)
	out := make(chan *parser.ParseResult)
	numWorkers := 10
	log.Println("Creating workers ...")
	for i := 0; i < numWorkers; i++ {
		CreateWorker(i, in, out)
	}
	log.Printf("Created %v workers", numWorkers)
	// non blocking send the seeds
	go func() {
		for _, seed := range seeds {
			in <- seed
		}
		log.Printf("Initial seeds are filled")
	}()

	var users []parser.UserProfile
	for {
		result := <-out
		if result == nil {
			continue
		}
		if result.Items != nil {
			users = append(users, result.Items...)
			log.Printf("We now have %v user profiles", len(users))
		}
		go func() {
			for _, r := range result.Requests {
				in <- r
			}
		}()
	}
}

// CreateWorker creates a worker that continuesly pulling new requests from the pool to do the work
func CreateWorker(id int, in chan parser.Request, out chan *parser.ParseResult) {
	go func() {
		for {
			r := <-in
			result, err := HandleRequest(r)
			if err != nil {
				log.Printf("Worker %d has an error %s", id, err.Error())
				continue
			}
			out <- result
		}
	}()
}
