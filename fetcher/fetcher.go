package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	ErrorResponseStatusCodeNotOK = errors.New("status code is not ok")
	rateLimiter                  = time.Tick(50 * time.Millisecond)
)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}
	url = strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s, received code is %v", ErrorResponseStatusCodeNotOK, res.StatusCode)
	}
	// when function ends, close the Closer.
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
